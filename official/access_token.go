package official

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"

	"github.com/dysodeng/wx/support/http"
)

// AccessToken 获取/刷新token
func (official *Official) AccessToken() (contracts.AccessToken, error) {
	return official.accessToken(false)
}

// accessToken 获取/刷新token
func (official *Official) accessToken(refresh bool) (contracts.AccessToken, error) {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.AuthorizerAccessToken(
			official.config.appId,
			official.config.authorizerRefreshToken,
			refresh,
			official.option.locker,
		)
	} else {
	cache:
		if !refresh && official.option.cache.IsExist(official.AccessTokenCacheKey()) {
			tokenString, err := official.option.cache.Get(official.AccessTokenCacheKey())
			if err == nil {
				var accessToken contracts.AccessToken
				err = json.Unmarshal([]byte(tokenString), &accessToken)
				if err == nil {
					return accessToken, nil
				}
			}
		}

		official.option.locker.Lock()
		defer func() {
			official.option.locker.Unlock()
		}()

		if official.option.cache.IsExist(official.AccessTokenCacheKey()) {
			goto cache
		}

		// 刷新access_token
		return official.refreshAccessToken()
	}
}

// refreshAccessToken 刷新access_token
func (official *Official) refreshAccessToken() (contracts.AccessToken, error) {
	apiUrl := fmt.Sprintf(
		"cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		official.config.appId,
		official.config.appSecret,
	)
	res, err := http.Get(apiUrl)
	if err != nil {
		return contracts.AccessToken{}, kernelError.New(0, err)
	}

	type accessToken struct {
		kernelError.ApiError
		contracts.AccessToken
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return contracts.AccessToken{}, kernelError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(contracts.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	})

	err = official.option.cache.Put(
		official.AccessTokenCacheKey(),
		string(tokenByte),
		time.Second*time.Duration(result.AccessToken.ExpiresIn-600),
	)
	if err != nil {
		return contracts.AccessToken{}, kernelError.New(0, err)
	}

	return contracts.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	}, nil
}

// AccessTokenCacheKey 获取access_token缓存key
func (official *Official) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", official.option.cacheKeyPrefix, "access_token", official.config.appId)
}
