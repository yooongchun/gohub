package common

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/controller/common"
)

var R = new(Router)

type Router struct{}

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/verify-code", func(group *ghttp.RouterGroup) {
		group.Bind(common.NewV1())
	})
}
