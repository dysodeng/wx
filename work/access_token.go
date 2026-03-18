package work

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/work/http"
)

func (w *Work) AccessToken() (contracts.AccessToken, error) {
	return w.accessToken(false)
}

func (w *Work) accessToken(refresh bool) (contracts.AccessToken, error) {
cache:
	if !refresh && w.option.cache.IsExist(w.AccessTokenCacheKey()) {
		tokenString, err := w.option.cache.Get(w.AccessTokenCacheKey())
		if err == nil {
			var accessToken contracts.AccessToken
			err = json.Unmarshal([]byte(tokenString), &accessToken)
			if err == nil {
				return accessToken, nil
			}
		}
	}

	w.option.locker.Lock()
	defer func() {
		w.option.locker.Unlock()
	}()

	if w.option.cache.IsExist(w.AccessTokenCacheKey()) {
		goto cache
	}

	// 刷新access_token
	return w.refreshAccessToken()
}

func (w *Work) refreshAccessToken() (contracts.AccessToken, error) {
	apiUrl := fmt.Sprintf(
		"cgi-bin/gettoken?corpid=%s&corpsecret=%s",
		w.config.corpId,
		w.config.secret,
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
		return contracts.AccessToken{}, kernelError.NewWithApiError(result.ApiError)
	}

	tokenByte, _ := json.Marshal(contracts.AccessToken{
		AccessToken: result.AccessToken.AccessToken,
		ExpiresIn:   result.AccessToken.ExpiresIn,
	})

	err = w.option.cache.Put(
		w.AccessTokenCacheKey(),
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

func (w *Work) AccessTokenCacheKey() string {
	return fmt.Sprintf("%s%s.%s", w.option.cacheKeyPrefix, "access_token", w.config.corpId)
}
