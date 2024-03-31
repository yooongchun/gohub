package system

import (
	"context"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) DeleteLoginLogList(ctx context.Context, req *v1.DeleteLoginLogListReq) (res *v1.DeleteLoginLogListRes, err error) {
	err = service.SysUser().Delete(ctx, req.Ids)
	return
}
