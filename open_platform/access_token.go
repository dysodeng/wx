package open_platform

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

const componentVerifyTicketCacheKey = "component_verify_ticket.%s"

// AccessToken 获取开放平台access_token
func (open *OpenPlatform) AccessToken() (kernel.AccessToken, error) {
	return open.accessToken(false)
}

func (open *OpenPlatform) accessToken(refresh bool) (kernel.AccessToken, error) {
	if !refresh && open.option.cache.IsExist(open.AccessTokenCacheKey()) {
		tokenString, err := open.option.cache.Get(open.AccessTokenCacheKey())
		if err == nil {
			var accessToken kernel.AccessToken
			err = json.Unmarshal([]byte(tokenString), &accessToken)
			if err == nil {
				return accessToken, nil
			}
		}
	}

	// 刷新access_token
	return open.refreshAccessToken()
}

func (open *OpenPlatform) refreshAccessToken() (kernel.AccessToken, error) {
	verifyTicket := open.getComponentVerifyTicket()
	res, err := http.PostJSON("cgi-bin/component/api_component_token", map[string]interface{}{
		"component_appid":         open.config.appId,
		"component_appsecret":     open.config.appSecret,
		"component_verify_ticket": verifyTicket,
	})
	if err != nil {
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	// 返回信息
	type accessToken struct {
		kernelError.ApiError
		ComponentAccessToken string `json:"component_access_token"`
		ExpiresIn            int64  `json:"expires_in"`
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return kernel.AccessToken{}, kernelError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(kernel.AccessToken{
		AccessToken: result.ComponentAccessToken,
		ExpiresIn:   result.ExpiresIn,
	})

	err = open.option.cache.Put(
		open.AccessTokenCacheKey(),
		string(tokenByte),
		time.Second*time.Duration(result.ExpiresIn-600), // 提前过期
	)
	if err != nil {
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	return kernel.AccessToken{
		AccessToken: result.ComponentAccessToken,
		ExpiresIn:   result.ExpiresIn,
	}, nil
}

func (open *OpenPlatform) getComponentVerifyTicket() string {
	cacheKey := open.option.cacheKeyPrefix + fmt.Sprintf(componentVerifyTicketCacheKey, open.config.appId)
	if open.option.cache.IsExist(cacheKey) {
		ticketString, err := open.option.cache.Get(cacheKey)
		if err == nil {
			return ticketString
		}
	}
	return ""
}

// AccessTokenCacheKey 获取开放平台access_token缓存key
func (open *OpenPlatform) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", open.option.cacheKeyPrefix, "component_access_token", open.config.appId)
}

// AuthorizerAccessToken 代公众账号获取access_token
func (open *OpenPlatform) AuthorizerAccessToken(appId, authorizerRefreshToken string, refresh bool) (kernel.AccessToken, error) {
	if !refresh && open.option.cache.IsExist(open.AuthorizerAccessTokenCacheKey(appId)) {
		tokenString, err := open.option.cache.Get(open.AuthorizerAccessTokenCacheKey(appId))
		if err == nil {
			var accessToken kernel.AccessToken
			err = json.Unmarshal([]byte(tokenString), &accessToken)
			if err == nil {
				return accessToken, nil
			}
		}
	}

	// 刷新公众账号access_token
	return open.refreshAuthorizerAccessToken(appId, authorizerRefreshToken)
}

func (open *OpenPlatform) refreshAuthorizerAccessToken(appId, authorizerRefreshToken string) (kernel.AccessToken, error) {
	componentAccessToken, err := open.AccessToken()
	if err != nil {
		return kernel.AccessToken{}, err
	}

	apiUrl := fmt.Sprintf("cgi-bin/component/api_authorizer_token?component_access_token=%s", componentAccessToken.AccessToken)
	res, err := http.PostJSON(apiUrl, map[string]interface{}{
		"component_appid":          open.config.appId,
		"authorizer_appid":         appId,
		"authorizer_refresh_token": authorizerRefreshToken,
	})
	if err != nil {
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	// 返回信息
	type accessToken struct {
		kernelError.ApiError
		AuthorizerAccessToken  string `json:"authorizer_access_token"`
		ExpiresIn              int64  `json:"expires_in"`
		AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return kernel.AccessToken{}, kernelError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(kernel.AccessToken{
		AccessToken: result.AuthorizerAccessToken,
		ExpiresIn:   result.ExpiresIn,
	})

	err = open.option.cache.Put(
		open.AuthorizerAccessTokenCacheKey(appId),
		string(tokenByte),
		time.Second*time.Duration(result.ExpiresIn-600), // 提前过期
	)
	if err != nil {
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	return kernel.AccessToken{
		AccessToken: result.AuthorizerRefreshToken,
		ExpiresIn:   result.ExpiresIn,
	}, nil
}

// AuthorizerAccessTokenCacheKey 公众账号access_token缓存key
func (open *OpenPlatform) AuthorizerAccessTokenCacheKey(appId string) string {
	return fmt.Sprintf("%s%s.%s.%s", open.option.cacheKeyPrefix, "authorizer_access_token", open.config.appId, appId)
}
