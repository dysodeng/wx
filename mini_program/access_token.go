package mini_program

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

func (mp *MiniProgram) AccessToken() (kernel.AccessToken, error) {
	return mp.accessToken(false)
}

// accessToken 获取/刷新token
func (mp *MiniProgram) accessToken(refresh bool) (kernel.AccessToken, error) {
	if mp.config.isOpenPlatform {
		return mp.config.authorizerAccount.AuthorizerAccessToken(
			mp.config.appId,
			mp.config.authorizerRefreshToken,
			refresh,
		)
	} else {
		if !refresh && mp.option.cache.IsExist(mp.AccessTokenCacheKey()) {
			tokenString, err := mp.option.cache.Get(mp.AccessTokenCacheKey())
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
	return mp.refreshAccessToken()
}

func (mp *MiniProgram) refreshAccessToken() (kernel.AccessToken, error) {
	apiUrl := fmt.Sprintf(
		"cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		mp.config.appId,
		mp.config.appSecret,
	)
	res, err := http.Get(apiUrl)
	if err != nil {
		return kernel.AccessToken{}, baseError.New(0, err)
	}

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

	err = mp.option.cache.Put(
		mp.AccessTokenCacheKey(),
		string(tokenByte),
		time.Second*time.Duration(result.AccessToken.ExpiresIn-600),
	)
	if err != nil {
		return kernel.AccessToken{}, baseError.New(0, err)
	}

	return kernel.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	}, nil
}

func (mp *MiniProgram) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", mp.option.cacheKeyPrefix, "access_token", mp.config.appId)
}

func (mp *MiniProgram) Cache() (cache.Cache, string) {
	return mp.option.cache, mp.option.cacheKeyPrefix
}
