// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"gohub/api/user/v1"
)

type IUserV1 interface {
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
	DeleteUser(ctx context.Context, req *v1.DeleteUserReq) (res *v1.DeleteUserRes, err error)
}
