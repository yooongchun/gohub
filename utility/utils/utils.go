package utils

import (
	"context"
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
