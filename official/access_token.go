package official

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/dysodeng/wx/kernel"
	kernelError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/http"
)

// AccessToken 获取/刷新token
func (official *Official) AccessToken() (kernel.AccessToken, error) {
	return official.accessToken(false)
}

// accessToken 获取/刷新token
func (official *Official) accessToken(refresh bool) (kernel.AccessToken, error) {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.AuthorizerAccessToken(
			official.config.appId,
			official.config.authorizerRefreshToken,
			refresh,
		)
	} else {
		if !refresh && official.option.cache.IsExist(official.AccessTokenCacheKey()) {
			tokenString, err := official.option.cache.Get(official.AccessTokenCacheKey())
			if err == nil {
				var accessToken kernel.AccessToken
				err = json.Unmarshal([]byte(tokenString), &accessToken)
				if err == nil {
					return accessToken, nil
				}
			}
		}
	}

	// 刷新access_token
	return official.refreshAccessToken()
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
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	type accessToken struct {
		kernelError.ApiError
		kernel.AccessToken
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return kernel.AccessToken{}, kernelError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(kernel.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	})

	err = official.option.cache.Put(
		official.AccessTokenCacheKey(),
		string(tokenByte),
		time.Second*time.Duration(result.AccessToken.ExpiresIn-600),
	)
	if err != nil {
		return kernel.AccessToken{}, kernelError.New(0, err)
	}

	return kernel.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	}, nil
}

// AccessTokenCacheKey 获取access_token缓存key
func (official *Official) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", official.option.cacheKeyPrefix, "access_token", official.config.appId)
}
