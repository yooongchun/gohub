package user

import (
	"context"
	v1 "gohub/api/user/v1"
	"gohub/internal/service"
)

func (c *ControllerV1) ClearOperateLog(ctx context.Context, req *v1.ClearOperateLogReq) (res *v1.ClearOperateLogRes, err error) {
	err = service.OperateLog().ClearLog(ctx)
	return
}
