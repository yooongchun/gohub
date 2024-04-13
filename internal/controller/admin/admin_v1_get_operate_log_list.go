package admin

import (
	"context"
	"gohub/internal/service"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) GetOperateLogList(ctx context.Context, req *v1.GetOperateLogListReq) (res *v1.GetOperateLogListRes, err error) {
	res, err = service.OperateLog().List(ctx, req)
	return
}
