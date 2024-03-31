package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error) {
	err = service.SysUser().Update(ctx, req, gmeta.Get(ctx, "id").Int64())
	return
}
