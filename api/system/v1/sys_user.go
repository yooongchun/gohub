package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/api/common/v1"
	"gohub/internal/model"
	"gohub/internal/model/entity"
)

// GetUserListReq 查询用户列表请求参数
type GetUserListReq struct {
	g.Meta   `path:"/users" tags:"用户管理" method:"get" summary:"用户列表"`
	Mobile   string `p:"mobile"`
	Status   string `p:"status"`
	KeyWords string `p:"keyWords"`
	Ids      []int  `p:"ids"`
	v1.PageReq
	v1.Author
}

// GetUserListRes 查询用户列表响应参数
type GetUserListRes struct {
	g.Meta   `mime:"application/json"`
	UserList []*model.LoginUserRes `json:"userList"`
	v1.ListRes
}

// CreateUserBaseReq 创建用户参数
type CreateUserBaseReq struct {
	Email    string `p:"email" v:"email#邮箱格式错误"` //邮箱
	NickName string `p:"nickName" v:"required#用户昵称不能为空"`
	Mobile   string `p:"mobile" v:"required|phone#手机号不能为空|手机号格式错误"`
	Remark   string `p:"remark"`
	Status   uint   `p:"status"`
}

// CreateUserReq 创建用户请求参数
type CreateUserReq struct {
	g.Meta `path:"/user" tags:"用户管理" method:"post" summary:"添加用户"`
	*CreateUserBaseReq
	UserName string `p:"userName" v:"required#用户账号不能为空"`
	Password string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	UserSalt string
}

type CreateUserRes struct {
}

// UpdateUserReq 修改用户请求参数
type UpdateUserReq struct {
	g.Meta `path:"/user/:id" tags:"用户管理" method:"put" summary:"修改用户"`
	*CreateUserBaseReq
}

type UpdateUserRes struct {
}

// GetUserOneReq 获取用户信息请求参数
type GetUserOneReq struct {
	g.Meta `path:"/user/:id" tags:"用户管理" method:"get" summary:"获取用户信息"`
}

// GetUserOneRes 获取用户信息响应参数
type GetUserOneRes struct {
	g.Meta `mime:"application/json"`
	User   *entity.SysUser `json:"user"`
}

// UpdateUserPwdReq 修改密码请求参数
type UpdateUserPwdReq struct {
	g.Meta   `path:"/user/:id/password" tags:"权限管理" method:"put" summary:"重置用户密码"`
	Password string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
}

type UpdateUserPwdRes struct {
}

// UpdateUserStatusReq 设置用户状态请求参数
type UpdateUserStatusReq struct {
	g.Meta     `path:"/user/status" tags:"用户管理" method:"put" summary:"设置用户状态"`
	Id         uint64 `p:"userId" v:"required#用户id不能为空"`
	UserStatus uint   `p:"status" v:"required#用户状态不能为空"`
}

type UpdateUserStatusRes struct {
}

// DeleteUserReq 删除用户请求参数
type DeleteUserReq struct {
	g.Meta `path:"/user" tags:"用户管理" method:"delete" summary:"删除用户"`
	Ids    []int `p:"ids"  v:"required#ids不能为空"`
}

type DeleteUserRes struct {
}
