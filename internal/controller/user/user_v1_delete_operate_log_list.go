package user

import (
	"context"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) DeleteOperateLogList(ctx context.Context, req *v1.DeleteOperateLogListReq) (res *v1.DeleteOperateLogListRes, err error) {
	err = service.OperateLog().DeleteByIds(ctx, req.OperateIds)
	return
}
