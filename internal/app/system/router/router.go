package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/app/system/controller"
	"gohub/internal/app/system/service"
)

var R Router

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Login) // Login
		// 登录验证拦截
		_ = service.GhToken().Middleware(group)
		// context 拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		//后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
		//group.Bind()
	})
}
