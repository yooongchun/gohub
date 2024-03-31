package token

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
	"gohub/internal/consts"
	"gohub/internal/model"
	"gohub/internal/service"
	"gohub/utility/errUtils"
)

type sToken struct {
	*gftoken.GfToken
}

func NewToken() *sToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		err = g.Cfg().MustGet(ctx, "ghToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	errUtils.ErrIfNotNil(ctx, err)
	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		)}
}

func init() {
	service.RegisterGhToken(NewToken())
}
