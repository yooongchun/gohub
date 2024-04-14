// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package common

import (
	"context"

	"gohub/api/common/v1"
)

type ICommonV1 interface {
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	PhoneLogin(ctx context.Context, req *v1.PhoneLoginReq) (res *v1.PhoneLoginRes, err error)
	EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error)
	UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
	UserRegisterUsingPhone(ctx context.Context, req *v1.UserRegisterUsingPhoneReq) (res *v1.UserRegisterUsingPhoneRes, err error)
	UserRegisterUsingEmail(ctx context.Context, req *v1.UserRegisterUsingEmailReq) (res *v1.UserRegisterUsingEmailRes, err error)
	GetVerifyCodeByCaptcha(ctx context.Context, req *v1.GetVerifyCodeByCaptchaReq) (res *v1.GetVerifyCodeByCaptchaRes, err error)
	GetVerifyCodeByEmail(ctx context.Context, req *v1.GetVerifyCodeByEmailReq) (res *v1.GetVerifyCodeByEmailRes, err error)
	GetVerifyCodeByPhone(ctx context.Context, req *v1.GetVerifyCodeByPhoneReq) (res *v1.GetVerifyCodeByPhoneRes, err error)
}
