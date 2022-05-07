package open_platform

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/support/cache"
)

// AccessToken 获取开放平台access_token
func (open *OpenPlatform) AccessToken(refresh bool) (base.AccessToken, error) {
	return base.AccessToken{}, nil
}

// AccessTokenKey 获取开放平台access_token缓存key
func (open *OpenPlatform) AccessTokenKey() string {
	return ""
}

// AuthorizerAccessToken 代第三方平台获取access_token
func (open *OpenPlatform) AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (base.AccessToken, error) {

	return base.AccessToken{}, nil
}

// AuthorizerAccessTokenKey 第三方平台access_token缓存key
func (open *OpenPlatform) AuthorizerAccessTokenKey() string {
	return ""
}

func (open *OpenPlatform) Cache() (cache.Cache, string) {
	return nil, ""
}
