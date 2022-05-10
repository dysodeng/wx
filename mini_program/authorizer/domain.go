package authorizer

import (
	"encoding/json"
	"fmt"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// Domain 小程序域名配置
type Domain struct {
	account contracts.AccountInterface
}

func NewDomain(account contracts.AccountInterface) *Domain {
	return &Domain{account: account}
}

// Modify 设置服务器域名
func (domain *Domain) Modify(data map[string][]string) (map[string]interface{}, error) {
	accountToken, err := domain.account.AccessToken()
	if err != nil {
		return nil, err
	}

	apiUrl := fmt.Sprintf("wxa/modify_domain?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, data)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	if result["errcode"] != 0 {
		return result, baseError.New(result["errcode"].(int64), errors.New(result["errmsg"].(string)))
	}

	return result, nil
}

// SetWebViewDomain 设置业务域名
func (domain *Domain) SetWebViewDomain(action string, domains ...string) error {
	accountToken, err := domain.account.AccessToken()
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("wxa/setwebviewdomain?access_token=%s", accountToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"action":        action,
		"webviewdomain": domains,
	})
	if err != nil {
		return err
	}

	var result baseError.WxApiError
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return nil
}
