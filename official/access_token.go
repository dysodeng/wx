package official

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dyaodeng/wx/base"
	baseError "github.com/dyaodeng/wx/base/error"
	"github.com/dyaodeng/wx/base/http"
	"time"
)

// AccessToken 获取/刷新token
func (official *Official) AccessToken(refresh bool) (base.AccessToken, error) {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccessToken.AccessToken(
			official.config.appId,
			official.config.authorizerRefreshToken,
			refresh,
		)
	} else {
		if !refresh && official.option.cache.IsExist(official.AccessTokenKey()) {
			tokenString, err := official.option.cache.Get(official.AccessTokenKey())
			if err == nil {
				if t, ok := tokenString.(string); ok {
					var accessToken base.AccessToken
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
	return fmt.Sprintf("%s%s", official.option.cacheKeyPrefix, "access_token")
}

// refreshAccessToken 刷新access_token
func (official *Official) refreshAccessToken() (base.AccessToken, error) {
	apiUrl := fmt.Sprintf(
		"cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		official.config.appId,
		official.config.appSecret,
	)
	res, err := http.Get(apiUrl)
	if err != nil {
		return base.AccessToken{}, baseError.New(0, err)
	}

	// 返回错误信息
	type accessToken struct {
		baseError.WxApiError
		base.AccessToken
	}
	var result accessToken
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return base.AccessToken{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	tokenByte, _ := json.Marshal(base.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	})

	err = official.option.cache.Put(
		official.AccessTokenKey(),
		string(tokenByte),
		time.Second*time.Duration(result.AccessToken.ExpiresIn-600), // 提前过期
	)
	if err != nil {
		return base.AccessToken{}, baseError.New(0, err)
	}

	return base.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	}, nil
}
