package base

// AccessToken access token
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// AccessTokenInterface 公众账号获取token接口
type AccessTokenInterface interface {
	AccessToken(refresh bool) (AccessToken, error)
	AccessTokenKey() string
}

// AuthorizerAccessTokenInterface 开放平台第三方授权公众账号获取token接口
type AuthorizerAccessTokenInterface interface {
	AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (AccessToken, error)
	AuthorizerAccessTokenKey() string
}
