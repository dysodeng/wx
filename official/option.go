package official

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/support/cache"
)

// config 公众号配置
type config struct {
	isOpenPlatform         bool
	appId                  string
	appSecret              string
	token                  string
	aesKey                 string
	authorizerRefreshToken string
	authorizerAccount      base.AuthorizerAccountInterface
}

// option 公众号选项
type option struct {
	cache          cache.Cache
	cacheKeyPrefix string
}

type Option func(*option)

// WithCache 设置缓存
func WithCache(cache cache.Cache) Option {
	return func(o *option) {
		o.cache = cache
	}
}

// WithCacheKeyPrefix 设置缓存key前缀
func WithCacheKeyPrefix(cacheKeyPrefix string) Option {
	return func(o *option) {
		o.cacheKeyPrefix = cacheKeyPrefix
	}
}
