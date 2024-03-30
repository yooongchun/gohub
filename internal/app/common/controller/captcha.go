package controller

import (
	"context"
	"gohub/api/v1/common"
	"gohub/internal/app/common/service"
)

var Captcha = captchaController{}

type captchaController struct{}

// Get 获取验证码
func (c *captchaController) Get(ctx context.Context, req *common.CaptchaReq) (res *common.CaptchaRes, err error) {
	var idKeyC, base64StringC string
	idKeyC, base64StringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &common.CaptchaRes{
		Key: idKeyC,
		Img: base64StringC,
	}
	return
}
