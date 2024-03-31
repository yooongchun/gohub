package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) GetOperateLogOne(ctx context.Context, req *v1.GetOperateLogOneReq) (res *v1.GetOperateLogOneRes, err error) {
	res = new(v1.GetOperateLogOneRes)
	res.SysOperateLogInfoRes, err = service.OperateLog().GetByOperateId(ctx, gmeta.Get(ctx, "operateId").Uint64())
	return
}
