package contracts

// AccountInterface 公众账号接口
type AccountInterface interface {
	AccessTokenInterface
	// IsOpenPlatform 是否为开放平台下的公众账号
	IsOpenPlatform() bool
	// AccountToken 公众账号token
	AccountToken() string
	// AccountAesKey 公众账号aes_key
	AccountAesKey() string
	// AccountAppId 公众账号appID
	AccountAppId() string
	// AccountAppSecret 公众账号app_secret
	AccountAppSecret() string
	// ComponentAppId 开放平台appID
	ComponentAppId() string
	// ComponentAccessToken 开放平台access_token
	ComponentAccessToken() string
}

// AuthorizerAccountInterface 开放平台第三方授权的公众账号接口
type AuthorizerAccountInterface interface {
	AuthorizerAccessTokenInterface
	// AuthorizerAccountToken 已授权公众号token
	AuthorizerAccountToken() string
	// AuthorizerAccountAesKey 已授权公众号aes_key
	AuthorizerAccountAesKey() string
	// ComponentAppId 开放平台appID
	ComponentAppId() string
	// ComponentAccessToken 开放平台access_token
	ComponentAccessToken() string
}
