package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"gohub/internal/model/entity"
	"gohub/internal/service"

	"gohub/api/system/v1"
)

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	var (
		userList []*entity.SysUser
	)
	res = new(v1.GetUserListRes)
	res.Total, userList, err = service.SysUser().List(ctx, req)
	if err != nil {
		return
	}
	err = gconv.Struct(userList, &res.UserList)
	return
}
