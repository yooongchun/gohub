package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmeta"
	"gohub/internal/model/entity"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) GetUserOne(ctx context.Context, req *v1.GetUserOneReq) (res *v1.GetUserOneRes, err error) {
	var user = new(entity.SysUser)
	user, err = service.SysUser().GetUserById(ctx, gmeta.Get(ctx, "id").Uint64())
	if err != nil {
		return
	}
	err = gconv.Struct(user, &res)
	return
}
