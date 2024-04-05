// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/os/gcache"
)

type (
	ICache interface {
		GetCache() *gcache.Cache
	}
)

var (
	localCache ICache
)

func Cache() ICache {
	if localCache == nil {
		panic("implement not found for interface ICache, forgot register?")
	}
	return localCache
}

func RegisterCache(i ICache) {
	localCache = i
}
