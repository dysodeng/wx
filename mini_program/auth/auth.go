package auth

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Auth 用户登录
// @see https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
type Auth struct {
	account contracts.AccountInterface
}

// Session 登录凭证
type Session struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionId    string `json:"unionid"`
}

func NewAuth(account contracts.AccountInterface) *Auth {
	return &Auth{account: account}
}

// Session 登录凭证校验
func (auth *Auth) Session(code string) (Session, error) {
	apiUrl := auth.getCodeUrl(code)
	res, err := http.Get(apiUrl)
	if err != nil {
		return Session{}, err
	}

	type sessionResult struct {
		kernelError.ApiError
		Session
	}
	var result sessionResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return Session{}, err
	}

	if result.ErrCode != 0 {
		return Session{}, errors.New(result.ErrMsg)
	}

	return result.Session, nil
}

func (auth *Auth) getCodeUrl(code string) string {
	if auth.account.IsOpenPlatform() {
		return fmt.Sprintf(
			"sns/component/jscode2session?appid=%s&js_code=%s&grant_type=authorization_code&component_appid=%s&component_access_token=%s",
			auth.account.AppId(),
			code,
			auth.account.ComponentAppId(),
			auth.account.ComponentAccessToken(),
		)
	}
	return fmt.Sprintf(
		"sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		auth.account.AppId(),
		auth.account.AppSecret(),
		code,
	)
}
