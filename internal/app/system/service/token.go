package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/tiger1103/gfast-token/gftoken" //#TODO: implement own token
)

type IGhToken interface {
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	Middleware(group *ghttp.RouterGroup) error
	ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	IsLogin(r *ghttp.Request) (b bool, failed *gftoken.AuthFailed)
	GetRequestToken(r *ghttp.Request) (token string)
	RemoveToken(ctx context.Context, token string) (err error)
	GetTokenData(ctx context.Context, token string) (tData *gftoken.TokenData, key string, err error)
}

var localGhToken IGhToken

func RegisterGhToken(i IGhToken) {
	localGhToken = i
}

func GhToken() IGhToken {
	if localGhToken == nil {
		panic("implement not found for interface IGhToken, forgot register?")
	}
	return localGhToken
}
