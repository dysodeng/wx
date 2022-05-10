package authorizer

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Tester 小程序成员管理
type Tester struct {
	account contracts.AccountInterface
}

func NewTester(account contracts.AccountInterface) *Tester {
	return &Tester{account: account}
}

// Bind 绑定体验者
// weChatId 微信号
func (tester *Tester) Bind(weChatId string) (string, error) {
	accountToken, err := tester.account.AccessToken()
	if err != nil {
		return "", err
	}

	apiUrl := fmt.Sprintf("wxa/bind_tester?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{"wechatid": weChatId})
	if err != nil {
		return "", baseError.New(0, err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return "", err
	}

	if result["errcode"] != 0 {
		return "", baseError.New(result["errcode"].(int64), errors.New(result["errmsg"].(string)))
	}

	return result["userstr"].(string), nil
}

// Unbind 解除绑定体验者
func (tester *Tester) Unbind(data map[string]string) error {
	accountToken, err := tester.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("wxa/unbind_tester?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, data)
	if err != nil {
		return baseError.New(0, err)
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err != nil {
		return err
	}

	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}

// GetMemberList 获取已绑定的体验者列表
func (tester Tester) GetMemberList() ([]map[string]interface{}, error) {
	accountToken, err := tester.account.AccessToken()
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("wxa/memberauth?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{"action": "get_experiencer"})
	if err != nil {
		return nil, baseError.New(0, err)
	}

	var result struct {
		baseError.WxApiError
		Members []map[string]interface{} `json:"members"`
	}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	if err == nil && result.ErrCode != 0 {
		return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.Members, nil
}
