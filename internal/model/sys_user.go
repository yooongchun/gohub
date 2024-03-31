package model

import "github.com/gogf/gf/v2/util/gmeta"

// LoginUserRes 登陆返回
type LoginUserRes struct {
	Id           uint64 `orm:"id,primary" json:"id"`               // 用户ID
	UserName     string `orm:"user_name,unique" json:"userName"`   // 用户名
	UserNickName string `orm:"user_nick_name" json:"userNickName"` // 用户昵称
	UserPassword string `orm:"user_password" json:"userPassword"`  // 用户密码
	UserSalt     string `orm:"user_salt" json:"userSalt"`          // 用户盐
	UserStatus   int    `orm:"user_status" json:"userStatus"`      // 用户状态;0:禁用,1:正常,2:未验证
	IsAdmin      int    `orm:"is_admin" json:"isAdmin"`            // 是否为管理员;0:否,1:是
	Avatar       string `orm:"avatar" json:"avatar"`               // 头像
}

// SysUserSimpleRes 用户简单信息
type SysUserSimpleRes struct {
	gmeta.Meta   `orm:"table:sys_user"`
	Id           uint64 `orm:"id"       json:"id"`                   //
	Avatar       string `orm:"avatar" json:"avatar"`                 // 头像
	UserName     string `orm:"user_name" json:"userName"`            // 用户名
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
}
