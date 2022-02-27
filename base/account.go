package base

// AccountInterface 公众账号接口
type AccountInterface interface {
	AccessTokenInterface
	AccountToken() string
	AccountAesKey() string
	AccountAppId() string
}

// AuthorizerAccountInterface 开放平台第三方授权的公众账号接口
type AuthorizerAccountInterface interface {
	AuthorizerAccessTokenInterface
	AuthorizerAccountToken() string
	AuthorizerAccountAesKey() string
}
