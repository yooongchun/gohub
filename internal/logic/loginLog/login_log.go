package loginLog

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/mssola/user_agent"
	"gohub/api/user/v1"
	"gohub/internal/consts"
	"gohub/internal/dao"
	"gohub/internal/model/do"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"gohub/utility/utils"
)

func init() {
	service.RegisterSysLoginLog(New())
}

func New() *sSysLoginLog {
	return &sSysLoginLog{
		Pool: grpool.New(100),
	}
}

type sSysLoginLog struct {
	Pool *grpool.Pool
}

func insertLoginLog(ctx context.Context, loginName, module, msg string, status uint) {
	//写入日志数据
	userAgent := utils.GetUserAgent(ctx)
	ua := user_agent.New(userAgent)
	browser, _ := ua.Browser()
	ip := utils.GetClientIp(ctx)

	loginData := &do.SysLoginLog{
		LoginName:     loginName,
		Ipaddr:        ip,
		LoginLocation: utils.GetCityByIp(ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        status,
		Msg:           msg,
		LoginTime:     gtime.Now(),
		Module:        module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	errUtils.ErrIfNotNil(ctx, err, "写入日志失败")
}

// Invoke 异步写入登录日志
func (s *sSysLoginLog) Invoke(ctx context.Context, loginName, module, msg string, status uint) {
	_ = s.Pool.Add(
		ctx,
		func(ctx context.Context) {
			//写入日志数据
			insertLoginLog(ctx, loginName, module, msg, status)
		},
	)
}

// List 获取登录日志列表
func (s *sSysLoginLog) List(ctx context.Context, req *v1.GetLoginLogListReq) (res *v1.GetLoginLogListRes, err error) {
	res = new(v1.GetLoginLogListRes)
	if req.PageNum == 0 {
		req.PageNum = 1
	}
	res.CurrentPage = req.PageNum
	if req.PageSize == 0 {
		req.PageSize = consts.PageSize
	}
	m := dao.SysLoginLog.Ctx(ctx)
	order := "info_id DESC"
	if req.LoginName != "" {
		m = m.Where("login_name like ?", "%"+req.LoginName+"%")
	}
	if req.Status != "" {
		m = m.Where("status", gconv.Int(req.Status))
	}
	if req.Ipaddr != "" {
		m = m.Where("ipaddr like ?", "%"+req.Ipaddr+"%")
	}
	if req.LoginLocation != "" {
		m = m.Where("login_location like ?", "%"+req.LoginLocation+"%")
	}
	if len(req.DateRange) != 0 {
		m = m.Where("login_time >=? AND login_time <=?", req.DateRange[0], req.DateRange[1])
	}
	if req.SortName != "" {
		if req.SortOrder != "" {
			order = req.SortName + " " + req.SortOrder
		} else {
			order = req.SortName + " DESC"
		}
	}
	err = g.Try(ctx, func(ctx context.Context) {
		res.Total, err = m.Count()
		errUtils.ErrIfNotNil(ctx, err, "获取日志失败")
		err = m.Page(req.PageNum, req.PageSize).Order(order).Scan(&res.List)
		errUtils.ErrIfNotNil(ctx, err, "获取日志数据失败")
	})
	return
}

// DeleteLoginLogByIds 删除登录日志
func (s *sSysLoginLog) DeleteLoginLogByIds(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysLoginLog.Ctx(ctx).Delete("info_id in (?)", ids)
		errUtils.ErrIfNotNil(ctx, err, "删除失败")
	})
	return
}

// ClearLoginLog 清空登录日志
func (s *sSysLoginLog) ClearLoginLog(ctx context.Context) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = g.DB().Ctx(ctx).Exec(ctx, "truncate "+dao.SysLoginLog.Table())
		errUtils.ErrIfNotNil(ctx, err, "清除失败")
	})
	return
}
