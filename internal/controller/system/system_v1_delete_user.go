package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}
