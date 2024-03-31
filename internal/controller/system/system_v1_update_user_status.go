package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) UpdateUserStatus(ctx context.Context, req *v1.UpdateUserStatusReq) (res *v1.UpdateUserStatusRes, err error) {
	err = service.SysUser().ChangeUserStatus(ctx, req, gmeta.Get(ctx, "id").Uint64())
	return
}
