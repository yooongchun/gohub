package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/internal/model/entity"
)

// GetUserOneReq 获取用户信息请求参数
type GetUserOneReq struct {
	g.Meta `path:"/user/:id" tags:"用户管理" method:"get" summary:"获取用户信息"`
}

// GetUserOneRes 获取用户信息响应参数
type GetUserOneRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *entity.SysUser `json:"user"`
}
