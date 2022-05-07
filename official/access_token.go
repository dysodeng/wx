package official

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel"
	baseError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/http"
)

// AccessToken 获取/刷新token
func (official *Official) AccessToken(refresh bool) (kernel.AccessToken, error) {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.AuthorizerAccessToken(
			official.config.appId,
			official.config.authorizerRefreshToken,
			refresh,
		)
	} else {
		if !refresh && official.option.cache.IsExist(official.AccessTokenKey()) {
			tokenString, err := official.option.cache.Get(official.AccessTokenKey())
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
	}

	// 刷新access_token
	return official.refreshAccessToken()
}

// AccessTokenKey 获取access_token缓存key
func (official *Official) AccessTokenKey() string {
	return fmt.Sprintf("%s%s:%s", official.option.cacheKeyPrefix, "access_token", official.config.appId)
}

// Cache 获取缓存实例
func (official *Official) Cache() (cache.Cache, string) {
	return official.option.cache, official.option.cacheKeyPrefix
}

// refreshAccessToken 刷新access_token
func (official *Official) refreshAccessToken() (kernel.AccessToken, error) {
	apiUrl := fmt.Sprintf(
		"cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		official.config.appId,
		official.config.appSecret,
	)
	res, err := http.Get(apiUrl)
	if err != nil {
		return kernel.AccessToken{}, baseError.New(0, err)
	}

	// 返回错误信息
	type accessToken struct {
		baseError.WxApiError
		kernel.AccessToken
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return kernel.AccessToken{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(kernel.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	})

	err = official.option.cache.Put(
		official.AccessTokenKey(),
		string(tokenByte),
		time.Second*time.Duration(result.AccessToken.ExpiresIn-600), // 提前过期
	)
	if err != nil {
		return kernel.AccessToken{}, baseError.New(0, err)
	}

	return kernel.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	}, nil
}
