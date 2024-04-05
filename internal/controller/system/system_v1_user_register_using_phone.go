package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"gohub/internal/dao"
	"gohub/internal/model/do"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"gohub/utility/utils"

	"gohub/api/system/v1"
)

func (c *ControllerV1) UserRegisterUsingPhone(ctx context.Context, req *v1.UserRegisterUsingPhoneReq) (res *v1.UserRegisterUsingPhoneRes, err error) {
	// 判断验证码是否正确
	err = checkVerifyCode(ctx, req.Mobile, req.VerifyCode)
	if err != nil {
		return
	}
	// 注册用户
	err = service.SysUser().UserExists(ctx, req.Mobile)
	if err != nil {
		return
	}
	salt := grand.S(10)
	req.Password = utils.EncryptPassword(req.Password, salt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.SysUser.Ctx(ctx).TX(tx).Insert(do.SysUser{
				UserName:     req.Username,
				UserPassword: req.Password,
				UserSalt:     salt,
				UserStatus:   1,
				Mobile:       req.Mobile,
				IsAdmin:      0,
			})
			errUtils.ErrIfNotNil(ctx, e, "注册失败，系统异常")
		})
		return err
	})
	return
}
