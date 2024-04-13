package admin

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/service"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	userId := ghttp.RequestFromCtx(ctx).Get("user_id").Uint64()
	err = service.SysUser().Delete(ctx, userId)
	return
}
