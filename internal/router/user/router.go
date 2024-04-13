package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/consts"
	"gohub/internal/controller/user"
	"gohub/internal/service"
	"gohub/utility/errUtils"
)

var R Router

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/", func(group *ghttp.RouterGroup) {
		// 登录验证拦截
		err := service.GhToken().Middleware(group)
		errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
		// 绑定后台路由
		group.Bind(user.NewV1())
	})
}
