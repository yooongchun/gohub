package sysUser

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/mssola/user_agent"
	"gohub/api/system/v1"
	"gohub/internal/consts"
	"gohub/internal/dao"
	"gohub/internal/model"
	"gohub/internal/model/do"
	"gohub/internal/model/entity"
	"gohub/internal/service"
	"gohub/utility/errUtils"
	"gohub/utility/utils"
)

func init() {
	service.RegisterSysUser(New())
}

type sSysUser struct {
}

func New() *sSysUser {
	return &sSysUser{}
}

func (s *sSysUser) NotCheckAuthAdminIds(ctx context.Context) *gset.Set {
	ids := g.Cfg().MustGet(ctx, "system.notCheckAuthAdminIds")
	if !g.IsNil(ids) {
		return gset.NewFrom(ids)
	}
	return gset.New()
}

func (s *sSysUser) GetAdminUserByUsernamePassword(ctx context.Context, req *v1.UserLoginReq) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		user, err = s.GetUserByUsername(ctx, req.Username)
		errUtils.ErrIfNotNil(ctx, err)
		errUtils.ErrIfValueNil(user, "账号密码错误")
		//验证密码
		if utils.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
			errUtils.ErrIfNotNil(ctx, gerror.New("账号密码错误"))
		}
		//账号状态
		if user.UserStatus == 0 {
			errUtils.ErrIfNotNil(ctx, gerror.New("账号已被冻结"))
		}
	})
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *sSysUser) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Fields(user).Where(dao.SysUser.Columns().UserName, userName).Scan(user)
		errUtils.ErrIfNotNil(ctx, err, "账号密码错误")
	})
	return
}

// GetUserById 通过用户名获取用户信息
func (s *sSysUser) GetUserById(ctx context.Context, id uint64) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Fields(user).WherePri(id).Scan(user)
		errUtils.ErrIfNotNil(ctx, err, "获取用户信息失败")
	})
	return
}

// LoginLog 记录登录日志
func (s *sSysUser) LoginLog(ctx context.Context, params *model.LoginLogParams) {
	ua := user_agent.New(params.UserAgent)
	browser, _ := ua.Browser()
	loginData := &do.SysLoginLog{
		LoginName:     params.Username,
		Ipaddr:        params.Ip,
		LoginLocation: utils.GetCityByIp(params.Ip),
		Browser:       browser,
		Os:            ua.OS(),
		Status:        params.Status,
		Msg:           params.Msg,
		LoginTime:     gtime.Now(),
		Module:        params.Module,
	}
	_, err := dao.SysLoginLog.Ctx(ctx).Insert(loginData)
	if err != nil {
		g.Log().Error(ctx, err)
	}
}

func (s *sSysUser) UpdateLoginInfo(ctx context.Context, id uint64, ip string) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Unscoped().Update(g.Map{
			dao.SysUser.Columns().LastLoginIp:   ip,
			dao.SysUser.Columns().LastLoginTime: gtime.Now(),
		})
		errUtils.ErrIfNotNil(ctx, err, "更新用户登录信息失败")
	})
	return
}

// List 用户列表
func (s *sSysUser) List(ctx context.Context, req *v1.GetUserListReq) (total interface{}, userList []*entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.KeyWords != "" {
			keyWords := "%" + req.KeyWords + "%"
			m = m.Where("user_name like ? or  user_nickname like ?", keyWords, keyWords)
		}
		if req.Status != "" {
			m = m.Where("user_status", gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where("mobile like ?", "%"+req.Mobile+"%")
		}
		if len(req.DateRange) > 0 {
			m = m.Where("created_at >=? AND created_at <=?", req.DateRange[0], req.DateRange[1])
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum <= 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		errUtils.ErrIfNotNil(ctx, err, "获取用户数据失败")
		err = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			Page(req.PageNum, req.PageSize).Order("id asc").Scan(&userList)
		errUtils.ErrIfNotNil(ctx, err, "获取用户列表失败")
	})
	return
}

func (s *sSysUser) Add(ctx context.Context, req *v1.CreateUserReq) (err error) {
	err = s.UserNameOrMobileExists(ctx, req.UserName, req.Mobile)
	if err != nil {
		return
	}
	req.UserSalt = grand.S(10)
	req.Password = utils.EncryptPassword(req.Password, req.UserSalt)
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, e := dao.SysUser.Ctx(ctx).TX(tx).Insert(do.SysUser{
				UserName:     req.UserName,
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserPassword: req.Password,
				UserSalt:     req.UserSalt,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Remark:       req.Remark,
			})
			errUtils.ErrIfNotNil(ctx, e, "添加用户失败")
		})
		return err
	})
	return
}

func (s *sSysUser) Update(ctx context.Context, req *v1.UpdateUserReq, id int64) (err error) {
	err = s.UserNameOrMobileExists(ctx, "", req.Mobile, id) //除自身之外手机号是否已存在
	if err != nil {
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(id).Update(do.SysUser{
				Mobile:       req.Mobile,
				UserNickname: req.NickName,
				UserStatus:   req.Status,
				UserEmail:    req.Email,
				Remark:       req.Remark,
			})
			errUtils.ErrIfNotNil(ctx, err, "修改用户信息失败")
		})
		return err
	})
	return
}

// UserNameOrMobileExists 用户名或手机号是否存在，如果传入id则需排除该id
func (s *sSysUser) UserNameOrMobileExists(ctx context.Context, userName, mobile string, id ...int64) error {
	user := (*entity.SysUser)(nil)
	err := g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if len(id) > 0 {
			m = m.Where(dao.SysUser.Columns().Id+" != ", id)
		}
		m = m.Where(fmt.Sprintf("%s='%s' OR %s='%s'",
			dao.SysUser.Columns().UserName,
			userName,
			dao.SysUser.Columns().Mobile,
			mobile))
		err := m.Limit(1).Scan(&user)
		errUtils.ErrIfNotNil(ctx, err, "获取用户信息失败")
		if user == nil {
			return
		}
		if user.UserName == userName {
			errUtils.ErrIfNotNil(ctx, gerror.New("用户名已存在"))
		}
		if user.Mobile == mobile {
			errUtils.ErrIfNotNil(ctx, gerror.New("手机号已存在"))
		}
	})
	return err
}

// GetUserInfo 获取编辑用户信息
func (s *sSysUser) GetUserInfo(ctx context.Context, id uint64) (res *v1.GetUserOneRes, err error) {
	res = new(v1.GetUserOneRes)
	err = g.Try(ctx, func(ctx context.Context) {
		//获取用户信息
		res.User, err = s.GetUserInfoById(ctx, id)
		errUtils.ErrIfNotNil(ctx, err)
	})
	return
}

// GetUserInfoById 通过Id获取用户信息
func (s *sSysUser) GetUserInfoById(ctx context.Context, id uint64, withPwd ...bool) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		if len(withPwd) > 0 && withPwd[0] {
			//用户用户信息
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).Scan(&user)
		} else {
			//用户用户信息
			err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id, id).
				FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).Scan(&user)
		}
		errUtils.ErrIfNotNil(ctx, err, "获取用户数据失败")
	})
	return
}

// ResetUserPwd 重置用户密码
func (s *sSysUser) ResetUserPwd(ctx context.Context, req *v1.UpdateUserPwdReq, id uint64) (err error) {
	salt := grand.S(10)
	password := utils.EncryptPassword(req.Password, salt)
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(g.Map{
			dao.SysUser.Columns().UserSalt:     salt,
			dao.SysUser.Columns().UserPassword: password,
		})
		errUtils.ErrIfNotNil(ctx, err, "重置用户密码失败")
	})
	return
}

func (s *sSysUser) ChangeUserStatus(ctx context.Context, req *v1.UpdateUserStatusReq, id uint64) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysUser.Ctx(ctx).WherePri(id).Update(do.SysUser{UserStatus: req.UserStatus})
		errUtils.ErrIfNotNil(ctx, err, "设置用户状态失败")
	})
	return
}

// Delete 删除用户
func (s *sSysUser) Delete(ctx context.Context, ids []int) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).Where(dao.SysUser.Columns().Id+" in(?)", ids).Delete()
			errUtils.ErrIfNotNil(ctx, err, "删除用户失败")
		})
		return err
	})
	return
}

// GetUsers 通过用户ids查询多个用户信息
func (s *sSysUser) GetUsers(ctx context.Context, ids []int) (users []*model.SysUserSimpleRes, err error) {
	if len(ids) == 0 {
		return
	}
	idsSet := gset.NewIntSetFrom(ids).Slice()
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Id+" in(?)", idsSet).
			Order(dao.SysUser.Columns().Id + " ASC").Scan(&users)
	})
	return
}
