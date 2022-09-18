package contracts

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
