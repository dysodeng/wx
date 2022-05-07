package contracts

import (
	"github.com/dysodeng/wx/kernel"
	"github.com/dysodeng/wx/support/cache"
)

// AccessTokenInterface 公众账号获取token接口
type AccessTokenInterface interface {
	// AccessToken 获取公众账号access_token
	AccessToken(refresh bool) (kernel.AccessToken, error)
	// AccessTokenKey 获取公众账号access_token缓存key
	AccessTokenKey() string
	// Cache 获取缓存实例
	Cache() (cache.Cache, string)
}

// AuthorizerAccessTokenInterface 开放平台第三方授权公众账号获取token接口
type AuthorizerAccessTokenInterface interface {
	// AuthorizerAccessToken 获取授权到开放平台的公众账号access_token
	AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (kernel.AccessToken, error)
	// AuthorizerAccessTokenKey 获取授权到开放平台的公众账号access_token缓存key
	AuthorizerAccessTokenKey() string
	// Cache 获取缓存实例
	Cache() (cache.Cache, string)
}
