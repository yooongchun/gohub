package admin

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/service"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) GetOperateLogOne(ctx context.Context, req *v1.GetOperateLogOneReq) (res *v1.GetOperateLogOneRes, err error) {
	res = &v1.GetOperateLogOneRes{}
	operateId := ghttp.RequestFromCtx(ctx).Get("operateId").Uint64()
	res.SysOperateLog, err = service.OperateLog().GetByOperateId(ctx, operateId)
	return
}
