package v1

import "github.com/gogf/gf/v2/frame/g"

type GetCaptchaOneReq struct {
	g.Meta `path:"/" tags:"验证码" method:"get" summary:"获取验证码"`
}

type GetCaptchaOneRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key"`
	Img    string `json:"img"`
}
