package common

import (
	"context"
	"gohub/internal/service"

	"gohub/api/common/v1"
)

func (c *ControllerV1) GetVerifyCodeByCaptcha(ctx context.Context, req *v1.GetVerifyCodeByCaptchaReq) (res *v1.GetVerifyCodeByCaptchaRes, err error) {
	var idKeyC, base64StringC string
	idKeyC, base64StringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &v1.GetVerifyCodeByCaptchaRes{
		Key: idKeyC,
		Img: base64StringC,
	}
	return
}
