package v1

import "github.com/gogf/gf/v2/frame/g"

type GetVerifyCodeByPhoneReq struct {
	g.Meta     `path:"/phone" tags:"手机验证码" method:"get" summary:"获取手机验证码"`
	Mobile     string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey" v:"required#验证码key不能为空"`
}

type GetVerifyCodeByPhoneRes struct {
	g.Meta `mime:"application/json"`
}
