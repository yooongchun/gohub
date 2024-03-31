package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "gohub/internal/app/common/router"
	commonService "gohub/internal/app/common/service"
	systemRouter "gohub/internal/app/system/router"
)

var R = new(Router)

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 跨域处理
		group.Middleware(commonService.Middleware().MiddlewareCORS)
		group.Middleware(ghttp.MiddlewareHandlerResponse)
		// 绑定后台理由
		systemRouter.R.BindController(ctx, group)
		//绑定公共路由
		commonRouter.R.BindController(ctx, group)
	})
}
