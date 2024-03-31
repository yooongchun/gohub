package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) CreateUser(ctx context.Context, req *v1.CreateUserReq) (res *v1.CreateUserRes, err error) {
	err = service.SysUser().Add(ctx, req)
	return
}
