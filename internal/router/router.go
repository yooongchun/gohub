package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/router/common"
	"gohub/internal/router/user"
	"gohub/internal/service"
)

var R = new(Router)

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		//跨域处理
		group.Middleware(service.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 错误处理中间件
		group.Middleware(service.Middleware().ErrorHandler)
		//绑定公共路由
		common.R.BindController(ctx, group)
		//绑定用户路由
		user.R.BindController(ctx, group)
		// 后台操作日志记录
		group.Hook("/*", ghttp.HookAfterOutput, service.OperateLog().OperationLog)
	})
}
