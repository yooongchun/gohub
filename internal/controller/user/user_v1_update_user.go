package user

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	userId := ghttp.RequestFromCtx(ctx).Get("id").Uint64()
	err = service.SysUser().Update(ctx, req, userId)
	return
}
