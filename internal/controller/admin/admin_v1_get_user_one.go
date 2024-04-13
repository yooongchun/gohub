package admin

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/service"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) GetUserOne(ctx context.Context, req *v1.GetUserOneReq) (res *v1.GetUserOneRes, err error) {
	userId := ghttp.RequestFromCtx(ctx).Get("id").Uint64()
	res = new(v1.GetUserOneRes)
	res.User, err = service.SysUser().GetUserById(ctx, userId)
	return
}
