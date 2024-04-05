// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	"gohub/api/system/v1"
)

type ISystemV1 interface {
	UserLogin(ctx context.Context, req *v1.UserLoginReq) (res *v1.UserLoginRes, err error)
	PhoneLogin(ctx context.Context, req *v1.PhoneLoginReq) (res *v1.PhoneLoginRes, err error)
	EmailLogin(ctx context.Context, req *v1.EmailLoginReq) (res *v1.EmailLoginRes, err error)
	UserLogout(ctx context.Context, req *v1.UserLogoutReq) (res *v1.UserLogoutRes, err error)
	UserRegisterUsingPhone(ctx context.Context, req *v1.UserRegisterUsingPhoneReq) (res *v1.UserRegisterUsingPhoneRes, err error)
	UserRegisterUsingEmail(ctx context.Context, req *v1.UserRegisterUsingEmailReq) (res *v1.UserRegisterUsingEmailRes, err error)
	GetLoginLogList(ctx context.Context, req *v1.GetLoginLogListReq) (res *v1.GetLoginLogListRes, err error)
	DeleteLoginLogList(ctx context.Context, req *v1.DeleteLoginLogListReq) (res *v1.DeleteLoginLogListRes, err error)
	ClearLoginLog(ctx context.Context, req *v1.ClearLoginLogReq) (res *v1.ClearLoginLogRes, err error)
	GetOperateLogList(ctx context.Context, req *v1.GetOperateLogListReq) (res *v1.GetOperateLogListRes, err error)
	GetOperateLogOne(ctx context.Context, req *v1.GetOperateLogOneReq) (res *v1.GetOperateLogOneRes, err error)
	DeleteOperateLogList(ctx context.Context, req *v1.DeleteOperateLogListReq) (res *v1.DeleteOperateLogListRes, err error)
	ClearOperateLog(ctx context.Context, req *v1.ClearOperateLogReq) (res *v1.ClearOperateLogRes, err error)
	GetUserList(ctx context.Context, req *v1.GetUserListReq) (res *v1.GetUserListRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
	GetUserOne(ctx context.Context, req *v1.GetUserOneReq) (res *v1.GetUserOneRes, err error)
	UpdateUserPwd(ctx context.Context, req *v1.UpdateUserPwdReq) (res *v1.UpdateUserPwdRes, err error)
	UpdateUserStatus(ctx context.Context, req *v1.UpdateUserStatusReq) (res *v1.UpdateUserStatusRes, err error)
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
}
