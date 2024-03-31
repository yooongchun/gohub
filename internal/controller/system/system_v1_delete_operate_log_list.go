package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) DeleteOperateLogList(ctx context.Context, req *v1.DeleteOperateLogListReq) (res *v1.DeleteOperateLogListRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperateIds)
	return
}
