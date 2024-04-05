package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"gohub/api/system/v1"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var resBase *v1.LoginResCommon
	err, resBase = LoginCommon(ctx, req.Username, req.VerifyCode, "username")
	if err != nil {
		return
	}
	err = gconv.Struct(resBase, res)
	return
}
