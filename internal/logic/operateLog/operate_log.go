package operateLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"gohub/api/admin/v1"
	"gohub/internal/consts"
	"gohub/internal/dao"
	"gohub/internal/model/do"
	"gohub/internal/model/entity"
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

// OperationLog 操作日志写入用于hook操作
func (s *sOperateLog) OperationLog(r *ghttp.Request) {
	s.Invoke(gctx.New(), r.GetMap(), 1)
}

// OperationLog 操作日志写入
func (s *sOperateLog) insertOperateLog(ctx context.Context, userName string, params g.Map, operatorType uint) {
	//写入日志数据
	r := ghttp.RequestFromCtx(ctx)
	ip := utils.GetClientIp(ctx)
	insertData := &do.SysOperateLog{
		Title:           "系统操作",
		Method:          r.URL.Path,
		RequestMethod:   r.Method,
		OperatorType:    operatorType,
		OperateName:     userName,
		OperateIp:       ip,
		OperateLocation: utils.GetCityByIp(ip),
		OperateTime:     gtime.Now(),
		OperateParam:    params,
	}
	rawQuery := r.URL.RawQuery
	if rawQuery != "" {
		rawQuery = "?" + rawQuery
	}
	insertData.OperateUrl = r.URL.Path + rawQuery
	_, err := dao.SysOperateLog.Ctx(ctx).Insert(insertData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

// Invoke 异步写入操作日志
func (s *sOperateLog) Invoke(ctx context.Context, params g.Map, operatorType uint) {
	userInfo := service.Context().GetLoginUser(ctx)
	if userInfo == nil {
		return
	}
	_ = s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		s.insertOperateLog(ctx, userInfo.UserName, params, operatorType)
	})
}

// List 操作日志列表
func (s *sOperateLog) List(ctx context.Context, req *v1.GetOperateLogListReq) (listRes *v1.GetOperateLogListRes, err error) {
	listRes = new(v1.GetOperateLogListRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysOperateLog.Ctx(ctx)
		if req.Title != "" {
			m = m.Where(dao.SysOperateLog.Columns().Title, req.Title)
		}
		if req.RequestMethod != "" {
			m = m.Where(dao.SysOperateLog.Columns().RequestMethod, req.RequestMethod)
		}
		if req.OperateName != "" {
			m = m.WhereLike(dao.SysOperateLog.Columns().OperateName, req.OperateName)
		}
		if len(req.DateRange) != 0 {
			m = m.WhereGTE(dao.SysOperateLog.Columns().OperateTime, req.DateRange[0]).
				WhereLTE(dao.SysOperateLog.Columns().OperateTime, req.DateRange[1])
		}
		listRes.Total, err = m.Count()
		errUtils.ErrIfNotNil(ctx, err, "获取总行数失败")
		if req.PageNum <= 0 {
			req.PageNum = 1
		}
		listRes.CurrentPage = req.PageNum
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.OrderBy == "asc" {
			m = m.OrderAsc(dao.SysOperateLog.Columns().OperateId)
		} else {
			m = m.OrderDesc(dao.SysOperateLog.Columns().OperateId)
		}
		var res []*entity.SysOperateLog
		err = m.Fields(v1.GetOperateLogListRes{}).Page(req.PageNum, req.PageSize).Scan(&res)
		errUtils.ErrIfNotNil(ctx, err, "获取数据失败")
	})
	return
}

// GetByOperateId 根据id获取操作日志
func (s *sOperateLog) GetByOperateId(ctx context.Context, operateId uint64) (res *entity.SysOperateLog, err error) {
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
