package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"gohub/api/system/v1"
)

func (c *ControllerV1) EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error) {
	var resBase *v1.LoginResCommon
	err, resBase = LoginCommon(ctx, req.Email, req.VerifyCode, "email")
	if err != nil {
		return
	}
	err = gconv.Struct(resBase, res)
	return
}
