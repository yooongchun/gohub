package common

import (
	"context"
	"gohub/internal/service"

	"gohub/api/common/v1"
)

func (c *ControllerV1) GetCaptchaOne(ctx context.Context, req *v1.GetCaptchaOneReq) (res *v1.GetCaptchaOneRes, err error) {
	var idKeyC, base64StringC string
	idKeyC, base64StringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &v1.GetCaptchaOneRes{
		Key: idKeyC,
		Img: base64StringC,
	}
	return
}
