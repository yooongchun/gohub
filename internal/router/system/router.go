package system

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/consts"
	"gohub/internal/controller/system"
	"gohub/internal/service"
	service2 "gohub/internal/service"
	"gohub/utility/errUtils"
)

var R Router

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/sys", func(group *ghttp.RouterGroup) {
		group.Bind(system.NewV1()) // Login
		// 登录验证拦截
		err := service2.GhToken().Middleware(group)
		errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
		// context 拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		//group.Bind()
	})
}
