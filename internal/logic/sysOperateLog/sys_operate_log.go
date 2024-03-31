package sysOperateLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"gohub/api/system/v1"
	"gohub/internal/consts"
	"gohub/internal/dao"
	"gohub/internal/model"
	"gohub/internal/model/do"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"gohub/utility/utils"
)

type sOperateLog struct {
	Pool *grpool.Pool
}

func init() {
	service.RegisterOperateLog(New())
}

func New() *sOperateLog {
	return &sOperateLog{
		Pool: grpool.New(100),
	}
}

// OperationLog 操作日志写入
func (s *sOperateLog) OperationLog(r *ghttp.Request) {
	userInfo := service.Context().GetLoginUser(r.GetCtx())
	if userInfo == nil {
		return
	}
	url := r.Request.URL //请求地址
	data := &model.SysOperateLogAdd{
		User:         userInfo,
		Url:          url,
		Params:       r.GetMap(),
		Method:       r.Method,
		ClientIp:     utils.GetClientIp(r.GetCtx()),
		OperatorType: 1,
	}
	s.Invoke(gctx.New(), data)
}

func (s *sOperateLog) Invoke(ctx context.Context, data *model.SysOperateLogAdd) {
	_ = s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		s.operationLogAdd(ctx, data)
	})
}

// OperationLogAdd 添加操作日志
func (s *sOperateLog) operationLogAdd(ctx context.Context, data *model.SysOperateLogAdd) {
	insertData := &do.SysOperateLog{
		Title:           "系统操作",
		Method:          data.Url.Path,
		RequestMethod:   data.Method,
		OperatorType:    data.OperatorType,
		OperateName:     data.User.UserName,
		OperateIp:       data.ClientIp,
		OperateLocation: utils.GetCityByIp(data.ClientIp),
		OperateTime:     gtime.Now(),
		OperateParam:    data.Params,
	}
	rawQuery := data.Url.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	insertData.OperateUrl = data.Url.Path + rawQuery
	_, err := dao.SysOperateLog.Ctx(ctx).Insert(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// List 操作日志列表
func (s *sOperateLog) List(ctx context.Context, req *v1.GetOperateLogListReq) (listRes *v1.GetOperateLogListRes, err error) {
	listRes = new(v1.GetOperateLogListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysOperateLog.Ctx(ctx)
		if req.Title != "" {
			m = m.Where(dao.SysOperateLog.Columns().Title+" = ?", req.Title)
		}
		if req.RequestMethod != "" {
			m = m.Where(dao.SysOperateLog.Columns().RequestMethod+" = ?", req.RequestMethod)
		}
		if req.OperateName != "" {
			m = m.Where(dao.SysOperateLog.Columns().OperateName+" like ?", "%"+req.OperateName+"%")
		}
		if len(req.DateRange) != 0 {
			m = m.Where("operate_time >=? AND operate_time <=?", req.DateRange[0], req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		errUtils.ErrIfNotNil(ctx, err, "获取总行数失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		order := "operate_id DESC"
		if req.OrderBy != "" {
			order = req.OrderBy
		}
		var res []*model.SysOperateLogInfoRes
		err = m.Fields(v1.GetOperateLogListRes{}).Page(req.PageNum, req.PageSize).Order(order).Scan(&res)
		errUtils.ErrIfNotNil(ctx, err, "获取数据失败")
		listRes.List = make([]*model.SysOperateLogListRes, len(res))
		for k, v := range res {
			listRes.List[k] = &model.SysOperateLogListRes{
				OperateId:       v.OperateId,
				Title:           v.Title,
				RequestMethod:   v.RequestMethod,
				OperateName:     v.OperateName,
				OperateUrl:      v.OperateUrl,
				OperateIp:       v.OperateIp,
				OperateLocation: v.OperateLocation,
				OperateParam:    v.OperateParam,
				OperateTime:     v.OperateTime,
			}
		}
	})
	return
}

// GetByOperateId 根据id获取操作日志
func (s *sOperateLog) GetByOperateId(ctx context.Context, operateId uint64) (res *model.SysOperateLogInfoRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysOperateLog.Ctx(ctx).WithAll().Where(dao.SysOperateLog.Columns().OperateId, operateId).Scan(&res)
		errUtils.ErrIfNotNil(ctx, err, "获取信息失败")
	})
	return
}

// DeleteByIds 根据id删除操作日志
func (s *sOperateLog) DeleteByIds(ctx context.Context, ids []uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysOperateLog.Ctx(ctx).Delete("operate_id in (?)", ids)
		errUtils.ErrIfNotNil(ctx, err, "删除失败")
	})
	return
}

// ClearLog 清除操作日志
func (s *sOperateLog) ClearLog(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = g.DB().Ctx(ctx).Exec(ctx, "truncate "+dao.SysOperateLog.Table())
		errUtils.ErrIfNotNil(ctx, err, "清除失败")
	})
	return
}
