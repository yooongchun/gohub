package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"gohub/internal/model/entity"
)

type UserLoginReq struct {
	g.Meta     `path:"/login/username" tags:"用户登陆" method:"post" summary:"账号密码登陆"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type LoginResCommon struct {
	g.Meta   `mime:"application/json"`
	UserInfo *entity.SysUser `json:"userInfo"`
	Token    string          `json:"token"`
}
type UserLoginRes struct {
	*LoginResCommon
}

type PhoneLoginReq struct {
	g.Meta     `path:"/login/phone" tags:"用户登陆" method:"post" summary:"手机号登陆"`
	Mobile     string `p:"phone" v:"required|phone#手机号不能为空|手机号不正确"`
	VerifyCode string `p:"verifyCode" v:"required|size:6#验证码不能为空|验证码长度为6位"`
}

type PhoneLoginRes struct {
	*LoginResCommon
}

type EmailLoginReq struct {
	g.Meta     `path:"/login/email" tags:"用户登陆" method:"post" summary:"邮箱登陆"`
	Email      string `p:"email" v:"required|email#邮箱不能为空|邮箱号无效"`
	VerifyCode string `p:"verifyCode" v:"required|size:6#验证码不能为空|验证码长度为6位"`
}

type EmailLoginRes struct {
	*LoginResCommon
}

type UserLogoutReq struct {
	g.Meta `path:"/logout" tags:"用户登录" method:"post" summary:"退出登录"`
	Author
}

type UserLogoutRes struct {
	g.Meta `mime:"application/json"`
}

type UserRegisterUsingPhoneReq struct {
	g.Meta     `path:"/register/phone" tags:"用户注册" method:"post" summary:"手机号注册"`
	Username   string `p:"username" v:"required|max-length:20#用户名不能为空|用户名长度不能超过20个字符"`
	Password   string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	VerifyCode string `p:"verifyCode" v:"required|size:6|integer#验证码不能为空|验证码长度为6位|验证码格式不正确"`
	Mobile     string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
}

type UserRegisterUsingEmailReq struct {
	g.Meta     `path:"/register/email" tags:"用户注册" method:"post" summary:"邮箱号注册"`
	Username   string `p:"username" v:"required|max-length:20#用户名不能为空|用户名长度不能超过20个字符"`
	Password   string `p:"password" v:"required|password#密码不能为空|密码以字母开头，只能包含字母、数字和下划线，长度在6~18之间"`
	VerifyCode string `p:"verifyCode" v:"required|size:6|integer#验证码不能为空|验证码长度为6位|验证码格式不正确"`
	Email      string `p:"email" v:"required|email#邮箱号不能为空|邮箱号格式不正确"`
}

type UserRegisterUsingPhoneRes struct {
	EmptyRes
}

type UserRegisterUsingEmailRes struct {
	EmptyRes
}
