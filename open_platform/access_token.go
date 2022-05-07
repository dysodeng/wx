package open_platform

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel"
	"github.com/dysodeng/wx/support/cache"
)

// AccessToken 获取开放平台access_token
func (open *OpenPlatform) AccessToken(refresh bool) (kernel.AccessToken, error) {
	if !refresh && open.option.cache.IsExist(open.AccessTokenKey()) {
		tokenString, err := open.option.cache.Get(open.AccessTokenKey())
		if err == nil {
			if t, ok := tokenString.(string); ok {
				var accessToken kernel.AccessToken
				err = json.Unmarshal([]byte(t), &accessToken)
				if err == nil {
					return accessToken, nil
				}
			}
		}
	}

	// 刷新access_token
	return open.refreshAccessToken()
}

func (open *OpenPlatform) refreshAccessToken() (kernel.AccessToken, error) {

	return kernel.AccessToken{}, nil
}

// AccessTokenKey 获取开放平台access_token缓存key
func (open *OpenPlatform) AccessTokenKey() string {
	return fmt.Sprintf("%s%s:%s", open.option.cacheKeyPrefix, "component_access_token", open.config.appId)
}

// AuthorizerAccessToken 代第三方平台获取access_token
func (open *OpenPlatform) AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (kernel.AccessToken, error) {

	return kernel.AccessToken{}, nil
}

// AuthorizerAccessTokenKey 第三方平台access_token缓存key
func (open *OpenPlatform) AuthorizerAccessTokenKey() string {
	return ""
}

func (open *OpenPlatform) Cache() (cache.Cache, string) {
	return open.option.cache, open.option.cacheKeyPrefix
}
