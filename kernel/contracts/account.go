package contracts

import "github.com/dysodeng/wx/support/cache"

// AccountInterface 公众账号接口
type AccountInterface interface {
	AccessTokenInterface
	// IsOpenPlatform 是否为开放平台下授权的公众账号
	IsOpenPlatform() bool
	// Token 获取公众账号token
	Token() string
	// AesKey 获取公众账号aes_key
	AesKey() string
	// AppId 获取公众账号appID
	AppId() string
	// AppSecret 获取公众账号app_secret
	AppSecret() string
	// ComponentAppId 获取开放平台appID
	ComponentAppId() string
	// ComponentAccessToken 获取开放平台access_token
	ComponentAccessToken() string
	// Cache 获取缓存实例
	Cache() (cache.Cache, string)
}

// AuthorizerInterface 开放平台-第三方平台接口(代公众账号发起接口调用方)
type AuthorizerInterface interface {
	AuthorizerAccessTokenInterface
	// AuthorizerAccountToken 获取已授权公众账号token
	AuthorizerAccountToken() string
	// AuthorizerAccountAesKey 获取已授权公众账号aes_key
	AuthorizerAccountAesKey() string
	// ComponentAppId 获取开放平台账号appID
	ComponentAppId() string
	// ComponentAccessToken 获取开放平台账号access_token
	ComponentAccessToken() string
}
