package user

import (
	"context"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) GetUserOne(ctx context.Context, req *v1.GetUserOneReq) (res *v1.GetUserOneRes, err error) {
	res.UserInfo, err = service.SysUser().GetUserById(ctx, gmeta.Get(ctx, "id").Uint64())
	if err != nil {
		return
	}
	return
}
