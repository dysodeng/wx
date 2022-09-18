package official

import (
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
)

// config 公众号配置
type config struct {
	isOpenPlatform         bool
	appId                  string
	appSecret              string
	token                  string
	aesKey                 string
	authorizerRefreshToken string
	authorizerAccount      contracts.AuthorizerInterface
}

// option 公众号选项
type option struct {
	cache          cache.Cache
	cacheKeyPrefix string
	locker         lock.Locker
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

// WithLocker 设置锁
func WithLocker(locker lock.Locker) Option {
	return func(o *option) {
		o.locker = locker
	}
}
