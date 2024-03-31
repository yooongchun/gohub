package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) ClearOperateLog(ctx context.Context, req *v1.ClearOperateLogReq) (res *v1.ClearOperateLogRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
