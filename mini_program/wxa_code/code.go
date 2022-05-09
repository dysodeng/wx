package wxa_code

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/dysodeng/wx/kernel/contracts"
	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

// WxaCode 小程序码
// @see https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/qr-code.html
// @see https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Business/qrcode.generate.html
type WxaCode struct {
	account contracts.AccountInterface
}

func NewWxaCode(account contracts.AccountInterface) *WxaCode {
	return &WxaCode{account: account}
}

// CreateQrCode 获取小程序二维码，适用于需要的码数量较少的业务场景
// 通过该接口生成的小程序码，永久有效，有数量限制
func (code *WxaCode) CreateQrCode(path string, opts map[string]interface{}) ([]byte, error) {
	if opts == nil {
		opts = make(map[string]interface{})
	}
	opts["path"] = path

	accountToken, err := code.account.AccessToken()
	if err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf("cgi-bin/wxaapp/createwxaqrcode?access_token=%s", accountToken.AccessToken)

	res, contentType, err := http.PostJSONWithRespContentType(apiUrl, opts)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(contentType, "application/json") {
		var result baseError.WxApiError
		err = json.Unmarshal(res, &result)
		if err == nil && result.ErrCode != 0 {
			return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
		}
	}

	if strings.HasPrefix(contentType, "image") {
		return res, nil
	}

	return nil, nil
}

// Get 获取小程序码，适用于需要的码数量较少的业务场景
// 通过该接口生成的小程序码，永久有效，有数量限制
func (code *WxaCode) Get(path string, opts map[string]interface{}) ([]byte, error) {
	if opts == nil {
		opts = make(map[string]interface{})
	}
	opts["path"] = path

	accountToken, err := code.account.AccessToken()
	if err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf("cgi-bin/wxaapp/getwxacode?access_token=%s", accountToken.AccessToken)

	res, contentType, err := http.PostJSONWithRespContentType(apiUrl, opts)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(contentType, "application/json") {
		var result baseError.WxApiError
		err = json.Unmarshal(res, &result)
		if err == nil && result.ErrCode != 0 {
			return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
		}
	}

	if strings.HasPrefix(contentType, "image") {
		return res, nil
	}

	return nil, nil
}

// GetUnlimited 获取小程序码，适用于需要的码数量极多的业务场景
// 通过该接口生成的小程序码，永久有效，数量暂无限制
func (code *WxaCode) GetUnlimited(scene string, opts map[string]interface{}) ([]byte, error) {
	if opts == nil {
		opts = make(map[string]interface{})
	}
	opts["scene"] = scene

	accountToken, err := code.account.AccessToken()
	if err != nil {
		return nil, err
	}
	apiUrl := fmt.Sprintf("cgi-bin/wxaapp/getwxacodeunlimit?access_token=%s", accountToken.AccessToken)

	res, contentType, err := http.PostJSONWithRespContentType(apiUrl, opts)
	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(contentType, "application/json") {
		var result baseError.WxApiError
		err = json.Unmarshal(res, &result)
		if err == nil && result.ErrCode != 0 {
			return nil, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
		}
	}

	if strings.HasPrefix(contentType, "image") {
		return res, nil
	}

	return nil, nil
}
