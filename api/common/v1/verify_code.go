package v1

import "github.com/gogf/gf/v2/frame/g"

// GetVerifyCodeByCaptchaReq 图形验证码
type GetVerifyCodeByCaptchaReq struct {
	g.Meta `path:"/captcha" tags:"验证码" method:"get" summary:"获取验证码"`
}

// GetVerifyCodeByCaptchaRes 图形验证码
type GetVerifyCodeByCaptchaRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}

// GetVerifyCodeByEmailReq 邮箱验证码
type GetVerifyCodeByEmailReq struct {
	g.Meta     `path:"/email" tags:"验证码" method:"get" summary:"获取邮箱验证码"`
	Email      string `p:"email" v:"required|email#邮箱不能为空|邮箱格式不正确"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

// GetVerifyCodeByEmailRes 邮箱验证码
type GetVerifyCodeByEmailRes struct {
	g.Meta `mime:"application/json"`
}

// GetVerifyCodeByPhoneReq 手机验证码
type GetVerifyCodeByPhoneReq struct {
	g.Meta     `path:"/phone" tags:"验证码" method:"get" summary:"获取手机验证码"`
	Mobile     string `p:"phone" v:"required|phone#手机号不能为空|手机号格式不正确"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey" v:"required#验证码key不能为空"`
}

// GetVerifyCodeByPhoneRes 手机验证码
type GetVerifyCodeByPhoneRes struct {
	EmptyRes
}
