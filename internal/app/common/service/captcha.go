package service

import "context"

type ICaptcha interface {
	GetVerifyImgString(ctx context.Context) (idKeyC string, base64StringC string, err error)
	VerifyCaptcha(idKey string, verifyValue string) bool
}

var localCaptcha ICaptcha

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}
func RegisterCaptcha(captcha ICaptcha) {
	localCaptcha = captcha
}
