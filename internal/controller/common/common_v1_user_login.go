package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"gohub/api/common/v1"
)

func (c *ControllerV1) UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error) {
	var resBase *v1.LoginResCommon
	err, resBase = LoginCommon(ctx, req.Username, req.VerifyKey, req.VerifyCode, "user")
	if err != nil {
		return
	}
	err = gconv.Struct(resBase, &res)
	return
}
