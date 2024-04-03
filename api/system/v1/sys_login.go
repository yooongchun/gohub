package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/api/common/v1"
	"gohub/internal/model"
)

type UserLoginReq struct {
	g.Meta     `path:"/login/username" tags:"登陆" method:"post" summary:"用户登陆"`
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
type PhoneLoginReq struct {
	g.Meta     `path:"/login/phone" tags:"手机验证码登陆" method:"post" summary:"用户登陆"`
	Mobile     string `p:"phone" v:"required#手机号不能为空"`
	VerifyCode string `p:"verifyCode" v:"required|size:6#手机验证码不能为空|手机验证码长度为6位"`
}
type PhoneLoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}
type UserLogoutReq struct {
	g.Meta `path:"/logout" tags:"登出" method:"post" summary:"用户登出"`
	v1.Author
}

type UserLogoutRes struct {
	g.Meta `mime:"application/json"`
}

type UserRegisterUsingPhoneReq struct {
	g.Meta     `path:"/register/phone" tags:"手机号注册" method:"post" summary:"用户注册"`
	Username   string `p:"username" v:"required|max-length:20#用户名不能为空|用户名长度不能超过20个字符"`
	Password   string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	VerifyCode string `p:"verifyCode" v:"required|size:6|integer#验证码不能为空|验证码长度为6位|验证码格式不正确"`
	Mobile     string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
}
type UserRegisterUsingEmailReq struct {
	g.Meta     `path:"/register/email" tags:"邮箱号注册" method:"post" summary:"用户注册"`
	Username   string `p:"username" v:"required|max-length:20#用户名不能为空|用户名长度不能超过20个字符"`
	Password   string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	VerifyCode string `p:"verifyCode" v:"required|size:6|integer#验证码不能为空|验证码长度为6位|验证码格式不正确"`
	Email      string `p:"email" v:"required|email#邮箱号不能为空|邮箱号格式不正确"`
}

type UserRegisterUsingPhoneRes struct {
	g.Meta `mime:"application/json"`
}
type UserRegisterUsingEmailRes struct {
	g.Meta `mime:"application/json"`
}
