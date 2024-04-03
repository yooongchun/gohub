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

func (c *ControllerV1) UserRegisterUsingEmail(ctx context.Context, req *v1.UserRegisterUsingEmailReq) (res *v1.UserRegisterUsingEmailRes, err error) {
	// 判断验证码是否正确
	err = checkVerifyCode(ctx, req.Email, req.VerifyCode)
	if err != nil {
		return
	}
	// 注册用户
	err = service.SysUser().UserExists(ctx, req.Username, "", req.Email)
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
				UserEmail:    req.Email,
				IsAdmin:      0,
			})
			errUtils.ErrIfNotNil(ctx, e, "添加用户失败")
		})
		return err
	})
	return
}
