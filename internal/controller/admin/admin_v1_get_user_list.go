package admin

import (
	"context"
	"gohub/internal/service"
	"gohub/utility/errUtils"

	"gohub/api/admin/v1"
)

func (c *ControllerV1) GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error) {
	res = new(v1.GetUserListRes)
	res.Total, res.UserList, err = service.SysUser().GetUserListPage(ctx, req)
	errUtils.ErrIfNotNil(ctx, err, "获取用户列表失败")
	return
}
