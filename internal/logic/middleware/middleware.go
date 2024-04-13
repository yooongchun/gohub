package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"gohub/internal/consts"
	"gohub/internal/model"
	"gohub/internal/service"
	"gohub/utility/responseUtils"
	"net/http"
	"strings"
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

// AdminRequired 权限判断处理中间件
func (s *sMiddleware) AdminRequired(r *ghttp.Request) {
	ctx := r.GetCtx()
	//获取登陆用户id
	ctxUser := service.Context().Get(ctx)
	if ctxUser.User.IsAdmin != 1 {
		responseUtils.FailJson(true, r, "您无权访问！")
	}
	r.Middleware.Next()
}

// LoginRequired 登录判断处理中间件
func (s *sMiddleware) LoginRequired(r *ghttp.Request) {
	ctx := r.GetCtx()
	//获取登陆用户id
	ctxUser := service.Context().Get(ctx)
	if ctxUser.User == nil {
		responseUtils.FailJson(true, r, "请先登录！")
	}
	r.Middleware.Next()
}

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorHandler 错误处理中间件
func (s *sMiddleware) ErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if err := r.GetError(); err != nil {
		r.Response.ClearBuffer()
		if strings.Contains(err.Error(), consts.InternalServerError) {
			r.Response.WriteStatus(gcode.New(http.StatusInternalServerError, consts.InternalServerError, nil).Code())
		} else {
			r.Response.WriteStatus(gcode.CodeOK.Code())
		}
		r.Response.WriteJson(APIResponse{Code: gerror.Code(err).Code(),
			Message: err.Error(),
			Data:    nil})
	}
}
