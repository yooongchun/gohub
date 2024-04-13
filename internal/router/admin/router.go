package admin

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/consts"
	"gohub/internal/controller/admin"
	"gohub/internal/service"
	"gohub/utility/errUtils"
)

var R Router

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/admin", func(group *ghttp.RouterGroup) {
		group.Bind(admin.NewV1()) // Login
		// 登录验证拦截
		err := service.GhToken().Middleware(group)
		errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
		// context 拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().AdminRequired)
	})
}
