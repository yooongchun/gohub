// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	ICaptcha interface {
		// GetVerifyImgString 生成验证码图片
		GetVerifyImgString(ctx context.Context) (idKeyC string, base64StringC string, err error)
		// VerifyCaptcha 验证验证码
		VerifyCaptcha(idKey string, verifyValue string) bool
	}
)

var (
	localCaptcha ICaptcha
)

func Captcha() ICaptcha {
	if localCaptcha == nil {
		panic("implement not found for interface ICaptcha, forgot register?")
	}
	return localCaptcha
}

func RegisterCaptcha(i ICaptcha) {
	localCaptcha = i
}
