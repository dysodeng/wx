package open

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Open 开放平台账号管理
type Open struct {
	account contracts.AccountInterface
}

func NewOpen(account contracts.AccountInterface) *Open {
	return &Open{account: account}
}

// Create 创建开放平台帐号并绑定公众号/小程序
// appId 公众号/小程序appId
func (o *Open) Create(appId string) (string, error) {
	accountToken, err := o.account.AccessToken()
	if err != nil {
		return "", err
	}

	apiUrl := fmt.Sprintf("cgi-bin/open/create?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"appid": appId,
	})
	if err != nil {
		return "", baseError.New(0, err)
	}

	var result openResult
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return "", baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.OpenAppid, nil
}

// Bind 将公众号/小程序绑定到开放平台帐号下
// appId 公众号/小程序appId
// openAppId 开放平台appId
func (o *Open) Bind(appId, openAppId string) error {
	accountToken, err := o.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("cgi-bin/open/bind?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"appid":      appId,
		"open_appid": openAppId,
	})
	if err != nil {
		return baseError.New(0, err)
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// Unbind 将公众号/小程序从开放平台帐号下解绑
// appId 公众号/小程序appId
// openAppId 开放平台appId
func (o *Open) Unbind(appId, openAppId string) error {
	accountToken, err := o.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("cgi-bin/open/unbind?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"appid":      appId,
		"open_appid": openAppId,
	})
	if err != nil {
		return baseError.New(0, err)
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// Get 获取公众号/小程序所绑定的开放平台帐号
// appId 公众号/小程序appId
func (o *Open) Get(appId string) (string, error) {
	accountToken, err := o.account.AccessToken()
	if err != nil {
		return "", err
	}

	apiUrl := fmt.Sprintf("cgi-bin/open/get?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"appid": appId,
	})
	if err != nil {
		return "", baseError.New(0, err)
	}

	var result openResult
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return "", baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.OpenAppid, nil
}

type openResult struct {
	baseError.WxApiError
	OpenAppid string `json:"open_appid"`
}
