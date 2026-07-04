package app

import (
	"github.com/goairix/wx/kernel/contracts"
	"github.com/goairix/wx/support/cache"
	"github.com/goairix/wx/support/lock"
)

type config struct {
	appId                  string
	appSecret              string
	token                  string
	aesKey                 string
	authorizerRefreshToken string
	authorizerAccount      contracts.AuthorizerInterface
}

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
