// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for table sys_user.
type SysUserDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SysUserColumns // columns contains all the column names of Table for convenient usage.
}

// SysUserColumns defines and stores column names for table sys_user.
type SysUserColumns struct {
	Id            string //
	UserName      string // 用户名
	Mobile        string // 手机号
	UserNickname  string // 用户昵称
	UserPassword  string // 登录密码;cmf_password加密
	UserSalt      string // 加密盐
	UserStatus    string // 用户状态;0:禁用,1:正常,2:未验证
	UserEmail     string // 用户登录邮箱
	Avatar        string // 用户头像
	Remark        string // 备注
	IsAdmin       string // 是否后台管理员 1 是  0   否
	Describe      string // 述信息
	LastLoginIp   string // 最后登录ip
	LastLoginTime string // 最后登录时间
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// sysUserColumns holds the columns for table sys_user.
var sysUserColumns = SysUserColumns{
	Id:            "id",
	UserName:      "user_name",
	Mobile:        "mobile",
	UserNickname:  "user_nickname",
	UserPassword:  "user_password",
	UserSalt:      "user_salt",
	UserStatus:    "user_status",
	UserEmail:     "user_email",
	Avatar:        "avatar",
	Remark:        "remark",
	IsAdmin:       "is_admin",
	Describe:      "describe",
	LastLoginIp:   "last_login_ip",
	LastLoginTime: "last_login_time",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao() *SysUserDao {
	return &SysUserDao{
		group:   "default",
		table:   "sys_user",
		columns: sysUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}