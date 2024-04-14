// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "gohub/api/user/v1"
	"gohub/internal/model/entity"
)

type (
	ISysUser interface {
		// GetUserById 通过用户名获取用户信息
		GetUserById(ctx context.Context, id uint64) (user *entity.SysUser, err error)
		// GetUserByUniqueKey 手机号，邮箱，用户名三者其一
		GetUserByUniqueKey(ctx context.Context, key string) (user *entity.SysUser, err error)
		// GetUserListPage 用户列表
		GetUserListPage(ctx context.Context, req *v1.GetUserListReq) (total interface{}, userList []*entity.SysUser, err error)
		// GetUserList 通过用户ids查询多个用户信息
		GetUserList(ctx context.Context, ids []int) (users []*entity.SysUser, err error)
		// GetUserListByQuery 通过用户ids查询多个用户信息
		GetUserListByQuery(ctx context.Context, query map[string]string) (users []*entity.SysUser, err error)
		// UpdateLoginInfo 更新登陆信息
		UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error)
		// Update 更新用户信息
		Update(ctx context.Context, req *v1.UpdateUserReq, id uint64) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, id uint64) (err error)
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
