package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/api/common/v1"
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
	UserList []*entity.SysUser `json:"userList"`
	v1.ListRes
}

// UpdateUserReq 修改用户请求参数
type UpdateUserReq struct {
	g.Meta    `path:"/user/:id" tags:"用户管理" method:"put" summary:"修改用户"`
	UserName  string `p:"userName"`
	Email     string `p:"email" v:"email#邮箱格式错误"` //邮箱
	NickName  string `p:"nickName" v:"用户昵称不能为空"`
	Mobile    string `p:"mobile" v:"phone#手机号格式错误"`
	Remark    string `p:"remark"`
	Status    uint   `p:"status"`
	Password  string `p:"password" v:"password#密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	Password2 string `p:"password" v:"password#密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
}

type UpdateUserRes struct {
	v1.EmptyRes
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

// DeleteUserReq 删除用户请求参数
type DeleteUserReq struct {
	g.Meta `path:"/user/:id" tags:"用户管理" method:"delete" summary:"删除用户"`
}

type DeleteUserRes struct {
	v1.EmptyRes
}

// GetUserInfoReq 获取用户信息请求参数
type GetUserInfoReq struct {
	g.Meta `path:"/user" tags:"用户管理" method:"get" summary:"个人信息"`
}

// GetUserInfoRes 获取用户信息响应参数
type GetUserInfoRes struct {
	g.Meta `mime:"application/json"`
	User   *entity.SysUser `json:"user"`
}
