package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) GetLoginLogList(ctx context.Context, req *v1.GetLoginLogListReq) (res *v1.GetLoginLogListRes, err error) {
	res, err = service.SysLoginLog().List(ctx, req)
	return
}
