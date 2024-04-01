package common

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gmode"
	"gohub/internal/service"

	"gohub/api/common/v1"
)

func (c *ControllerV1) GetVerifyCodeByEmail(ctx context.Context, req *v1.GetVerifyCodeByEmailReq) (res *v1.GetVerifyCodeByEmailRes, err error) {
	//判断验证码是否正确
	if !gmode.IsDevelop() {
		if !service.Captcha().VerifyCaptcha(req.VerifyKey, req.VerifyCode) {
			err = gerror.New("验证码输入错误")
			return
		}
	}
	err = service.Mail().Send(ctx, req.Email, "", "")
	return
}
