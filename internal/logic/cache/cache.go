package cache

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"gohub/internal/service"
)

func init() {
	service.RegisterCache(NewCache())
}

type sCache struct {
	cache *gcache.Cache
}

func NewCache() *sCache {
	// Create redis client object.
	cache := gcache.New()
	// Create redis cache adapter and set it to cache object.
	cache.SetAdapter(gcache.NewAdapterRedis(g.Redis("cache")))

	return &sCache{cache: cache}
}

func (s *sCache) GetCache() *gcache.Cache {
	return s.cache
}
