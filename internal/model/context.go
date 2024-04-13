package model

import "gohub/internal/model/entity"

type Context struct {
	User *ContextUser // User in context.
}

type ContextUser struct {
	*entity.SysUser
}
