package utils

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// GetClientIp 获取客户端IP
func GetClientIp(ctx context.Context) string {
	return g.RequestFromCtx(ctx).GetClientIp()
}

// GetUserAgent 获取客户端UserAgent
func GetUserAgent(ctx context.Context) string {
	return g.RequestFromCtx(ctx).Header.Get("User-Agent")
}

// EncryptPassword 密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

// GetCityByIp 获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "::1" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(context.TODO(), url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.Get("code").Int() == 0 {
		city := fmt.Sprintf("%s %s", json.Get("pro").String(), json.Get("city").String())
		return city
	} else {
		return ""
	}
}

// GetConfig 获取配置
func GetConfig(ctx context.Context, paramName string) string {
	return g.Cfg().MustGet(ctx, paramName).String()
}

// GetConfigInt 获取配置
func GetConfigInt(ctx context.Context, paramName string) int {
	return g.Cfg().MustGet(ctx, paramName).Int()
}
