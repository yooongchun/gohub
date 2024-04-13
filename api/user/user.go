// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"gohub/api/user/v1"
)

type IUserV1 interface {
	GetUserOne(ctx context.Context, req *v1.GetUserOneReq) (res *v1.GetUserOneRes, err error)
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	PhoneLogin(ctx context.Context, req *v1.PhoneLoginReq) (res *v1.PhoneLoginRes, err error)
	EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error)
	UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
	UserRegisterUsingPhone(ctx context.Context, req *v1.UserRegisterUsingPhoneReq) (res *v1.UserRegisterUsingPhoneRes, err error)
	UserRegisterUsingEmail(ctx context.Context, req *v1.UserRegisterUsingEmailReq) (res *v1.UserRegisterUsingEmailRes, err error)
}
