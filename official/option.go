package official

import (
	"github.com/dyaodeng/wx/base"
	"github.com/dyaodeng/wx/base/cache"
)

type config struct {
	isOpenPlatform         bool
	appId                  string
	appSecret              string
	token                  string
	aesKey                 string
	authorizerRefreshToken string
	authorizerAccessToken  base.AuthorizerAccessTokenInterface
}

type Config func(*config)

// WithOfficial 公众号
func WithOfficial(appId, appSecret, token, aesKey string) Config {
	return func(cfg *config) {
		cfg.isOpenPlatform = false
		cfg.appId = appId
		cfg.appSecret = appSecret
		cfg.token = token
		cfg.aesKey = aesKey
	}
}

// WithOpenPlatform 开放平台代公众号调用接口
func WithOpenPlatform(appId, authorizerRefreshToken string, authorizerAccessToken base.AuthorizerAccessTokenInterface) Config {
	return func(cfg *config) {
		cfg.isOpenPlatform = true
		cfg.appId = appId
		cfg.authorizerRefreshToken = authorizerRefreshToken
		cfg.authorizerAccessToken = authorizerAccessToken
	}
}

type option struct {
	cache          cache.Cache
	cacheKeyPrefix string
}

const DefaultCacheKeyPrefix = "dy.wx.official.cache."

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
