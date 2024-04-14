package common

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"

	"gohub/api/common/v1"
)

func (c *ControllerV1) PhoneLogin(ctx context.Context, req *v1.PhoneLoginReq) (res *v1.PhoneLoginRes, err error) {
	var resBase *v1.LoginResCommon
	err, resBase = LoginCommon(ctx, req.Mobile, "", req.VerifyCode, "mobile")
	if err != nil {
		return
	}
	err = gconv.Struct(resBase, res)
	return
}
