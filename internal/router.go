package internal

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
)

var R = new(Router)

type Router struct{}

func (r *Router) BindController(ctx *context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 跨域处理
		group.Middleware()
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		//绑定后台路由

		//绑定公公路由
		//自动绑定自定义模块
	})
}
