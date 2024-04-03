// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"gohub/utility/errUtils"
	"gohub/utility/utils"
)

func checkVerifyCode(ctx context.Context, key, code string) (err error) {
	redis := g.Redis()
	// 判断验证码是否正确
	cacheKeyPrefix := utils.GetConfig(ctx, "verifyCode.cacheKeyPrefix")
	var verifyCode *gvar.Var
	cacheKey := fmt.Sprintf("%s%s", cacheKeyPrefix, key)
	verifyCode, err = redis.Get(ctx, cacheKey)
	errUtils.ErrIfNotNil(ctx, err, "服务器内部错误")
	if verifyCode == nil {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "验证码已过期")
		return
	}
	if verifyCode.String() != code {
		err = gerror.NewCode(gcode.CodeInvalidParameter, "验证码输入错误")
		return
	}
	// 删除已验证的验证码
	_, err = redis.Del(ctx, cacheKey)
	errUtils.ErrIfNotNil(ctx, err, "服务器内部错误")
	return
}
