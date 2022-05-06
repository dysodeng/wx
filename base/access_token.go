package base

// AccessToken access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// AccessTokenInterface 公众账号获取token接口
type AccessTokenInterface interface {
	// AccessToken 获取公众账号access_token
	AccessToken(refresh bool) (AccessToken, error)
	// AccessTokenKey 获取公众账号access_token缓存key
	AccessTokenKey() string
}

// AuthorizerAccessTokenInterface 开放平台第三方授权公众账号获取token接口
type AuthorizerAccessTokenInterface interface {
	// AuthorizerAccessToken 获取授权到开放平台的公众账号access_token
	AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (AccessToken, error)
	// AuthorizerAccessTokenKey 获取授权到开放平台的公众账号access_token缓存key
	AuthorizerAccessTokenKey() string
}
