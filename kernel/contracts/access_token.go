package contracts

import (
	"github.com/dysodeng/wx/kernel"
	"github.com/dysodeng/wx/support/cache"
)

// AccessTokenInterface 公众账号获取token接口
type AccessTokenInterface interface {
	// AccessToken 获取公众账号access_token
	AccessToken() (kernel.AccessToken, error)
	// AccessTokenCacheKey 获取公众账号access_token缓存key
	AccessTokenCacheKey() string
	// Cache 获取缓存实例
	Cache() (cache.Cache, string)
}

// AuthorizerAccessTokenInterface 开放平台代公众账号获取token接口
type AuthorizerAccessTokenInterface interface {
	// AuthorizerAccessToken 代公众账号获取access_token
	AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (kernel.AccessToken, error)
	// AuthorizerAccessTokenCacheKey 公众账号access_token缓存key
	AuthorizerAccessTokenCacheKey(appId string) string
	// Cache 获取缓存实例
	Cache() (cache.Cache, string)
}
