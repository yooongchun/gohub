package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"gohub/internal/model"
	"gohub/internal/service"
	"gohub/utility/responseUtils"
)

func init() {
	service.RegisterMiddleware(NewMiddleware())
}

func NewMiddleware() *sMiddleware {
	return &sMiddleware{}
}

type sMiddleware struct {
}

func (s *sMiddleware) MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	// Allow all domains.
	corsOptions.AllowDomain = []string{"*"}
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// Ctx 自定义上下文对象
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	ctx := r.GetCtx()
	// 初始化登录用户信息
	data, err := service.GhToken().ParseToken(r)
	if err != nil {
		// 执行下一步请求逻辑
		r.Middleware.Next()
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(ctx, err)
			// 执行下一步请求逻辑
			r.Middleware.Next()
		}
		service.Context().Init(r, context)
	}
	// 执行下一步请求逻辑
	r.Middleware.Next()
}

// Auth 权限判断处理中间件
func (s *sMiddleware) Auth(r *ghttp.Request) {
	ctx := r.GetCtx()
	//获取登陆用户id
	adminId := service.Context().GetUserId(ctx)
	//获取无需验证权限的用户id
	tagSuperAdmin := false
	service.SysUser().NotCheckAuthAdminIds(ctx).Iterator(func(v interface{}) bool {
		if gconv.Uint64(v) == adminId {
			tagSuperAdmin = true
			return false
		}
		return true
	})
	if tagSuperAdmin {
		r.Middleware.Next()
		//不要再往后面执行
		return
	}
	responseUtils.FailJson(true, r, "无权访问")
}
