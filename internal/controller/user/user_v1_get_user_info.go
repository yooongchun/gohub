package user

import (
	"context"
	"gohub/internal/service"

	"gohub/api/user/v1"
)

func (c *ControllerV1) GetUserInfo(ctx context.Context, req *v1.GetUserInfoReq) (res *v1.GetUserInfoRes, err error) {
	//获取登陆用户id
	res = &v1.GetUserInfoRes{}
	res.User = service.Context().GetLoginUser(ctx).SysUser
	return
}
