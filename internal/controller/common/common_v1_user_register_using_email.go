package common

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"gohub/internal/consts"
	"gohub/internal/dao"
	"gohub/internal/model/do"
	"gohub/internal/model/entity"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"gohub/utility/utils"

	"gohub/api/common/v1"
)

func (c *ControllerV1) UserRegisterUsingEmail(ctx context.Context, req *v1.UserRegisterUsingEmailReq) (res *v1.UserRegisterUsingEmailRes, err error) {
	// 判断用户是否已存在
	var user = &entity.SysUser{}
	user, err = service.SysUser().GetUserByUniqueKey(ctx, req.Email)
	errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
	if user != nil {
		err = gerror.New("邮箱号已存在")
		return
	}
	user, err = service.SysUser().GetUserByUniqueKey(ctx, req.Username)
	errUtils.ErrIfNotNil(ctx, err, consts.InternalServerError)
	if user != nil {
		err = gerror.New("账号已存在")
		return
	}
	// 判断验证码是否正确
	err = checkVerifyCode(ctx, req.Email, req.VerifyCode)
	if err != nil {
		return
	}
	// 注册用户
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
			errUtils.ErrIfNotNil(ctx, e, "注册失败，系统异常")
		})
		return err
	})
	return
}
