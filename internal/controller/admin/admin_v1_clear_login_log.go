package admin

import (
	"context"
	"gohub/internal/service"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) ClearLoginLog(ctx context.Context, req *v1.ClearLoginLogReq) (res *v1.ClearLoginLogRes, err error) {
	err = service.SysLoginLog().ClearLoginLog(ctx)
	return
}
