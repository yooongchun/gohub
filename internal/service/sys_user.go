// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gohub/api/system/v1"
	"gohub/internal/model"
	"gohub/internal/model/entity"

	"github.com/gogf/gf/v2/container/gset"
)

type (
	ISysUser interface {
		NotCheckAuthAdminIds(ctx context.Context) *gset.Set
		GetAdminUserByUsernamePassword(ctx context.Context, req *v1.UserLoginReq) (user *model.LoginUserRes, err error)
		GetAdminUserByMobile(ctx context.Context, mobile string) (user *model.LoginUserRes, err error)
		// GetUserByUsername 通过用户名获取用户信息
		GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error)
		// GetUserById 通过用户名获取用户信息
		GetUserById(ctx context.Context, id uint64) (user *entity.SysUser, err error)
		// LoginLog 记录登录日志
		LoginLog(ctx context.Context, params *model.LoginLogParams)
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		// List 用户列表
		List(ctx context.Context, req *v1.GetUserListReq) (total interface{}, userList []*entity.SysUser, err error)
		Add(ctx context.Context, req *v1.CreateUserReq) (err error)
		Update(ctx context.Context, req *v1.UpdateUserReq, id int64) (err error)
		// UserNameOrMobileExists 用户名或手机号是否存在，如果传入id则需排除该id
		UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error
		// UserExists 用户是否存在，根据用户名，手机号，email三者其一判断
		UserExists(ctx context.Context, userName, mobile, email string) (err error)
		// GetUserInfo 获取编辑用户信息
		GetUserInfo(ctx context.Context, id uint64) (res *v1.GetUserOneRes, err error)
		// GetUserInfoById 通过Id获取用户信息
		GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error)
		// ResetUserPwd 重置用户密码
		ResetUserPwd(ctx context.Context, req *v1.UpdateUserPwdReq, id uint64) (err error)
		ChangeUserStatus(ctx context.Context, req *v1.UpdateUserStatusReq, id uint64) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, ids []int) (err error)
		// GetUsers 通过用户ids查询多个用户信息
		GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimpleRes, err error)
	}
)

var (
	localSysUser ISysUser
)

func SysUser() ISysUser {
	if localSysUser == nil {
		panic("implement not found for interface ISysUser, forgot register?")
	}
	return localSysUser
}

func RegisterSysUser(i ISysUser) {
	localSysUser = i
}
