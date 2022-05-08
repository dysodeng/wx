package open_platform

import (
	"encoding/json"
	"fmt"

	baseError "github.com/dysodeng/wx/kernel/error"
	"github.com/dysodeng/wx/support/http"
	"github.com/pkg/errors"
)

/**
 * 公众账号授权
 */

// AuthType 授权类型
type AuthType uint8

const (
	AuthOfficial    = 1 // 授权页仅展示公众号
	AuthMiniProgram = 2 // 授权页仅展示小程序
	AuthAll         = 3 // 授权页展示小程序与公众号
)

// PreAuthCode 预授权码
type PreAuthCode struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int64  `json:"expires_in"`
}

// AuthorizationInfo 授权信息
type AuthorizationInfo struct {
	AuthorizationInfo AuthorizationInfoBody `json:"authorization_info"`
}
type AuthorizationInfoBody struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int64      `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfo               []FuncInfo `json:"func_info,omitempty"`
}
type FuncInfo struct {
	FuncScopeCategory struct {
		Id int `json:"id"`
	} `json:"funcscope_category,omitempty"`
	ConfirmInfo struct {
		NeedConfirm    int `json:"need_confirm"`
		AlreadyConfirm int `json:"already_confirm"`
		CanConfirm     int `json:"can_confirm"`
	} `json:"confirm_info,omitempty"`
}

// AuthorizerInfo 授权账号详情信息
type AuthorizerInfo struct {
	AuthorizationInfo AuthorizationInfoBody `json:"authorization_info"`

	// 授权账号详情信息
	AuthorizerInfo struct {
		// 通用字段
		Nickname        string `json:"nick_name"`
		HeadImg         string `json:"head_img"`
		ServiceTypeInfo struct {
			Id int `json:"id"`
		} `json:"service_type_info"`
		VerifyTypeInfo struct {
			Id int `json:"id"`
		} `json:"verify_type_info"`
		Username      string         `json:"user_name"`
		PrincipalName string         `json:"principal_name"`
		BusinessInfo  map[string]int `json:"business_info"`
		Alias         string         `json:"alias"`
		QrcodeUrl     string         `json:"qrcode_url"`
		AccountStatus int            `json:"account_status"`

		// 小程序独有字段
		Idc             int             `json:"idc"`
		Signature       string          `json:"signature"`
		RegisterType    int             `json:"register_type"`
		BasicConfig     map[string]bool `json:"basic_config"`
		MiniProgramInfo struct {
			Network     map[string][]string `json:"network"`
			Categories  []Categories        `json:"categories"`
			VisitStatus int                 `json:"visit_status"`
		} `json:"MiniProgramInfo"`
	} `json:"authorizer_info"`
}
type Categories struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

// PreAuthorizationUrl 获取公众账号授权链接
func (open *OpenPlatform) PreAuthorizationUrl(callbackUrl string, authType AuthType) (string, error) {
	preAuthCode, err := open.getPreAuthCode()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"https://mp.weixin.qq.com/cgi-bin/componentloginpage?component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d",
		open.config.appId,
		preAuthCode.PreAuthCode,
		callbackUrl,
		authType,
	), nil
}

// MobilePreAuthorizationUrl 获取移动端公众账号授权链接
func (open *OpenPlatform) MobilePreAuthorizationUrl(callbackUrl string, authType AuthType) (string, error) {
	preAuthCode, err := open.getPreAuthCode()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"https://open.weixin.qq.com/wxaopen/safe/bindcomponent?action=%s&no_scan=1&component_appid=%s&pre_auth_code=%s&redirect_uri=%s&auth_type=%d#wechat_redirect",
		"bindcomponent",
		open.config.appId,
		preAuthCode.PreAuthCode,
		callbackUrl,
		authType,
	), nil
}

// getPreAuthCode 获取预授权码
func (open *OpenPlatform) getPreAuthCode() (PreAuthCode, error) {
	accessToken, err := open.AccessToken(false)
	if err != nil {
		return PreAuthCode{}, err
	}

	apiUrl := fmt.Sprintf("cgi-bin/component/api_create_preauthcode?component_access_token=%s", accessToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"component_appid": open.config.appId,
	})
	if err != nil {
		return PreAuthCode{}, baseError.New(0, err)
	}

	// 返回信息
	type preAuthCode struct {
		baseError.WxApiError
		PreAuthCode
	}
	var result preAuthCode
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return PreAuthCode{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.PreAuthCode, nil
}

// AuthorizationInfo 使用授权码获取授权信息
func (open *OpenPlatform) AuthorizationInfo(authCode string) (AuthorizationInfo, error) {
	accessToken, err := open.AccessToken(false)
	if err != nil {
		return AuthorizationInfo{}, err
	}

	apiUrl := fmt.Sprintf("cgi-bin/component/api_query_auth?component_access_token=%s", accessToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"component_appid":    open.config.appId,
		"authorization_code": authCode,
	})
	if err != nil {
		return AuthorizationInfo{}, baseError.New(0, err)
	}

	// 返回信息
	type info struct {
		baseError.WxApiError
		AuthorizationInfo
	}
	var result info
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return AuthorizationInfo{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.AuthorizationInfo, nil
}

// AuthorizerInfo 获取授权帐号详情
func (open *OpenPlatform) AuthorizerInfo(appId string) (AuthorizerInfo, error) {
	accessToken, err := open.AccessToken(false)
	if err != nil {
		return AuthorizerInfo{}, err
	}

	apiUrl := fmt.Sprintf("cgi-bin/component/api_get_authorizer_info?component_access_token==%s", accessToken.AccessToken)
	res, err := http.PostJson(apiUrl, map[string]interface{}{
		"component_appid":  open.config.appId,
		"authorizer_appid": appId,
	})
	if err != nil {
		return AuthorizerInfo{}, baseError.New(0, err)
	}

	// 返回信息
	type info struct {
		baseError.WxApiError
		AuthorizerInfo
	}
	var result info
	err = json.Unmarshal(res, &result)
	if err == nil && result.ErrCode != 0 {
		return AuthorizerInfo{}, baseError.New(result.ErrCode, errors.New(result.ErrMsg))
	}

	return result.AuthorizerInfo, nil
}
