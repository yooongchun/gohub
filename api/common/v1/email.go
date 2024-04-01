package v1

import "github.com/gogf/gf/v2/frame/g"

type GetVerifyCodeByEmailReq struct {
	g.Meta     `path:"/email" tags:"邮箱验证码" method:"get" summary:"获取邮箱验证码"`
	Email      string `p:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type GetVerifyCodeByEmailRes struct {
	g.Meta `mime:"application/json"`
}
