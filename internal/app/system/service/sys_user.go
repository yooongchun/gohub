package service

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"gohub/api/v1/system"
	"gohub/internal/app/system/model"
	"gohub/internal/app/system/model/entity"
)

type ISysUser interface {
	NotCheckAuthAdminIds(ctx context.Context) *gset.Set
	GetAdminUserByUsernamePassword(ctx context.Context, req *system.UserLoginReq) (user *model.LoginUserRes, err error)
	GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error)
	GetUserById(ctx context.Context, id uint64) (user *model.LoginUserRes, err error)
	LoginLog(ctx context.Context, params *model.LoginLogParams)
	UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
	List(ctx context.Context, req *system.UserSearchReq) (total interface{}, userList []*entity.SysUser, err error)
	Add(ctx context.Context, req *system.UserAddReq) (err error)
	Edit(ctx context.Context, req *system.UserEditReq) (err error)
	UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error
	GetEditUser(ctx context.Context, id uint64) (res *system.UserGetEditRes, err error)
	GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error)
	ResetUserPwd(ctx context.Context, req *system.UserResetPwdReq) (err error)
	ChangeUserStatus(ctx context.Context, req *system.UserStatusReq) (err error)
	Delete(ctx context.Context, ids []int) (err error)
	GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimpleRes, err error)
}

var localSysUser ISysUser

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
