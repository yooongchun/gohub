package user

import (
	"context"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) DeleteLoginLogList(ctx context.Context, req *v1.DeleteLoginLogListReq) (res *v1.DeleteLoginLogListRes, err error) {
	err = service.SysLoginLog().DeleteLoginLogByIds(ctx, req.Ids)
	return
}
