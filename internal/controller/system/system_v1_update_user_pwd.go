package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) UpdateUserPwd(ctx context.Context, req *v1.UpdateUserPwdReq) (res *v1.UpdateUserPwdRes, err error) {
	err = service.SysUser().ResetUserPwd(ctx, req, gmeta.Get(ctx, "id").Uint64())
	return
}
