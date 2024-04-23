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
		v1 := user.NewV1()
		// 登录验证拦截
		err := service.GhToken().Middleware(group)
		errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
		// Context 拦截，注入登录用户信息
		group.Middleware(service.Middleware().Ctx)
		group.Bind(v1.GetUserInfo)
		// 需要管理员
		group.Group("/admin", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware().Ctx, service.Middleware().AdminRequired)
			group.Bind(v1.GetLoginLogList)
			group.Bind(v1.DeleteLoginLogList)
			group.Bind(v1.ClearLoginLog)
			group.Bind(v1.GetOperateLogList)
			group.Bind(v1.GetOperateLogOne)
			group.Bind(v1.DeleteOperateLogList)
			group.Bind(v1.ClearOperateLog)
			group.Bind(v1.GetUserList)
			group.Bind(v1.UpdateUser)
			group.Bind(v1.GetUserOne)
			group.Bind(v1.DeleteUser)
		})
	})
}
