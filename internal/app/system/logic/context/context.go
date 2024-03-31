package context

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"gohub/internal/app/system/consts"
	"gohub/internal/app/system/model"
	"gohub/internal/app/system/service"
)

func init() {
	service.RegisterContext(NewContext())
}

type sContext struct{}

func NewContext() *sContext {
	return &sContext{}
}

// Init 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.CtxKey, customCtx)
}

// Get 获得上下文变量，如果没有设置，那么返回nil
func (s *sContext) Get(ctx context.Context) *model.Context {
	value := ctx.Value(consts.CtxKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetUser 将上下文信息设置到上下文请求中，注意是完整覆盖
func (s *sContext) SetUser(ctx context.Context, ctxUser *model.ContextUser) {
	s.Get(ctx).User = ctxUser
}

// GetLoginUser 获取当前登陆用户信息
func (s *sContext) GetLoginUser(ctx context.Context) *model.ContextUser {
	_context := s.Get(ctx)
	if _context == nil {
		return nil
	}
	return _context.User
}

// GetUserId 获取当前登录用户id
func (s *sContext) GetUserId(ctx context.Context) uint64 {
	user := s.GetLoginUser(ctx)
	if user != nil {
		return user.Id
	}
	return 0
}
