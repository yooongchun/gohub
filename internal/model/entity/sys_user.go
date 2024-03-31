// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id            uint64      `json:"id"            description:""`
	UserName      string      `json:"userName"      description:"用户名"`
	Mobile        string      `json:"mobile"        description:"手机号"`
	UserNickname  string      `json:"userNickname"  description:"用户昵称"`
	UserPassword  string      `json:"userPassword"  description:"登录密码;cmf_password加密"`
	UserSalt      string      `json:"userSalt"      description:"加密盐"`
	UserStatus    uint        `json:"userStatus"    description:"用户状态;0:禁用,1:正常,2:未验证"`
	UserEmail     string      `json:"userEmail"     description:"用户登录邮箱"`
	Avatar        string      `json:"avatar"        description:"用户头像"`
	Remark        string      `json:"remark"        description:"备注"`
	IsAdmin       int         `json:"isAdmin"       description:"是否后台管理员 1 是  0   否"`
	Describe      string      `json:"describe"      description:"述信息"`
	LastLoginIp   string      `json:"lastLoginIp"   description:"最后登录ip"`
	LastLoginTime *gtime.Time `json:"lastLoginTime" description:"最后登录时间"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
}
