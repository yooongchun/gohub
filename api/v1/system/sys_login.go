package system

import (
	"github.com/gogf/gf/v2/frame/g"
	commonAPI "gohub/api/v1/common"
	"gohub/internal/app/system/model"
)

type UserLoginReq struct {
	g.Meta     `path:"/login" tags:"登陆" method:"post" summary:"用户登陆"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}
type UserLoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type UserLogoutReq struct {
	g.Meta `path:"/logout" tags:"登出" method:"post" summary:"用户登出"`
	commonAPI.Author
}

type UserLogoutRes struct {
}