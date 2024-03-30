package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/app/common/controller"
)

var R = new(Router)

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/pub", func(group *ghttp.RouterGroup) {
		group.Group("/captcha", func(group *ghttp.RouterGroup) {
			group.Bind(controller.Captcha)
		})
	})
}
