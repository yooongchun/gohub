package sysUser

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"gohub/api/user/v1"
	"gohub/internal/consts"
	"gohub/internal/dao"
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

// GetUserById 通过用户名获取用户信息
func (s *sSysUser) GetUserById(ctx context.Context, id uint64) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			WherePri(id).Scan(&user)
		errUtils.ErrIfNotNil(ctx, err, "获取用户信息失败")
	})
	return
}

// GetUserByUniqueKey 手机号，邮箱，用户名三者其一
func (s *sSysUser) GetUserByUniqueKey(ctx context.Context, key string) (user *entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).
			FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			Where(dao.SysUser.Columns().UserName, key).
			WhereOr(dao.SysUser.Columns().Mobile, key).
			WhereOr(dao.SysUser.Columns().UserEmail).Scan(&user)
		errUtils.ErrIfNotNil(ctx, err)
		//账号状态
		if user != nil && user.UserStatus == 0 {
			errUtils.ErrIfNotNil(ctx, gerror.New("账号已被冻结"))
		}
	})
	return
}

// GetUserListPage 用户列表
func (s *sSysUser) GetUserListPage(ctx context.Context, req *v1.GetUserListReq) (total interface{}, userList []*entity.SysUser, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		if req.KeyWords != "" {
			m = m.WhereLike(dao.SysUser.Columns().UserName, req.KeyWords).WhereOrLike(dao.SysUser.Columns().UserNickname, req.KeyWords)
		}
		if req.Status != "" {
			m = m.Where(dao.SysUser.Columns().UserStatus, gconv.Int(req.Status))
		}
		if req.Mobile != "" {
			m = m.Where(dao.SysUser.Columns().Mobile, req.Mobile)
		}
		if len(req.DateRange) > 0 {
			m = m.WhereGTE(dao.SysUser.Columns().CreatedAt, req.DateRange[0]).WhereLTE(dao.SysUser.Columns().CreatedAt, req.DateRange[1])
		}
		if req.PageSize <= 0 {
			req.PageSize = consts.PageSize
		}
		if req.PageNum <= 0 {
			req.PageNum = 1
		}
		total, err = m.Count()
		errUtils.ErrIfNotNil(ctx, err, "获取用户数据失败")
		err = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			Page(req.PageNum, req.PageSize).OrderAsc(dao.SysUser.Columns().Id).Scan(&userList)
		errUtils.ErrIfNotNil(ctx, err, "获取用户列表失败")
	})
	return
}

// GetUserList 通过用户ids查询多个用户信息
func (s *sSysUser) GetUserList(ctx context.Context, ids []int) (users []*entity.SysUser, err error) {
	if len(ids) == 0 {
		return
	}
	idsSet := gset.NewIntSetFrom(ids).Slice()
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysUser.Ctx(ctx).FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			WhereIn(dao.SysUser.Columns().Id, idsSet).
			OrderAsc(dao.SysUser.Columns().Id).Scan(&users)
	})
	return
}

// GetUserListByQuery 通过用户ids查询多个用户信息
func (s *sSysUser) GetUserListByQuery(ctx context.Context, query map[string]string) (users []*entity.SysUser, err error) {
	if len(query) == 0 {
		return
	}
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysUser.Ctx(ctx)
		for key, val := range query {
			m = m.Where(key, val)
		}
		err = m.FieldsEx(dao.SysUser.Columns().UserPassword, dao.SysUser.Columns().UserSalt).
			OrderAsc(dao.SysUser.Columns().Id).Scan(&users)
	})
	return
}

// UpdateLoginInfo 更新登陆信息
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

// Update 更新用户信息
func (s *sSysUser) Update(ctx context.Context, req *v1.UpdateUserReq, id uint64) (err error) {
	if req.Password != "" && req.Password != req.Password2 {
		err = gerror.New("两次密码不一致")
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
			if req.Password != "" {
				salt := grand.S(10)
				password := utils.EncryptPassword(req.Password, salt)
				_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(id).Update(do.SysUser{
					UserPassword: password,
					UserSalt:     salt,
				})
				errUtils.ErrIfNotNil(ctx, err, "修改用户信息失败")
			}
		})
		return err
	})
	return
}

// Delete 删除用户
func (s *sSysUser) Delete(ctx context.Context, id uint64) (err error) {
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		err = g.Try(ctx, func(ctx context.Context) {
			_, err = dao.SysUser.Ctx(ctx).TX(tx).WherePri(id).Delete()
			errUtils.ErrIfNotNil(ctx, err, "删除用户失败")
		})
		return err
	})
	return
}
