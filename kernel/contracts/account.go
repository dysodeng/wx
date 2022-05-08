package contracts

// AccountInterface 公众账号接口
type AccountInterface interface {
	AccessTokenInterface
	// IsOpenPlatform 是否为开放平台下的公众账号
	IsOpenPlatform() bool
	// Token 公众账号token
	Token() string
	// AesKey 公众账号aes_key
	AesKey() string
	// AppId 公众账号appID
	AppId() string
	// AppSecret 公众账号app_secret
	AppSecret() string
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
