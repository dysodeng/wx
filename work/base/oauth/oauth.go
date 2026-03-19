package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/dysodeng/wx/kernel/contracts"
	kernelError "github.com/dysodeng/wx/kernel/error"
	workHttp "github.com/dysodeng/wx/work/http"
)

const oauthBaseUrl = "https://open.weixin.qq.com/connect/oauth2/authorize"

// OAuth 企业微信网页授权登录
type OAuth struct {
	account     contracts.AccountInterface
	scope       string
	redirectUrl string
	state       string
	agentId     string
}

func New(account contracts.AccountInterface) *OAuth {
	return &OAuth{account: account, state: "state", scope: "snsapi_base"}
}

// WithScope 设置授权作用域
func (auth *OAuth) WithScope(scope string) *OAuth {
	auth.scope = scope
	return auth
}

// WithRedirectUrl 设置回调URL
func (auth *OAuth) WithRedirectUrl(redirectUrl string) *OAuth {
	auth.redirectUrl = redirectUrl
	return auth
}

// WithState 设置state参数
func (auth *OAuth) WithState(state string) *OAuth {
	auth.state = state
	return auth
}

// WithAgentId 设置agentid
func (auth *OAuth) WithAgentId(agentId string) *OAuth {
	auth.agentId = agentId
	return auth
}

func (auth *OAuth) buildAuthUrl() string {
	authUrl := fmt.Sprintf(
		"%s?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s",
		oauthBaseUrl,
		auth.account.AppId(),
		url.QueryEscape(auth.redirectUrl),
		auth.scope,
		auth.state,
	)
	if auth.agentId != "" {
		authUrl += fmt.Sprintf("&agentid=%s", auth.agentId)
	}
	authUrl += "#wechat_redirect"
	return authUrl
}

// AuthUrl 构造授权链接
func (auth *OAuth) AuthUrl() string {
	return auth.buildAuthUrl()
}

// Redirect 302重定向到授权页
func (auth *OAuth) Redirect(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, auth.buildAuthUrl(), http.StatusFound)
}

// UserFromCode 通过code获取用户身份
func (auth *OAuth) UserFromCode(code string) (*UserIdentity, error) {
	accessToken, err := auth.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/auth/getuserinfo?access_token=%s&code=%s", accessToken.AccessToken, code)
	res, err := workHttp.Get(apiUrl)
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result userIdentityResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UserIdentity, nil
}

// GetUserDetail 获取用户敏感信息
func (auth *OAuth) GetUserDetail(userTicket string) (*UserDetail, error) {
	accessToken, err := auth.account.AccessToken()
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	apiUrl := fmt.Sprintf("cgi-bin/auth/getuserdetail?access_token=%s", accessToken.AccessToken)
	res, err := workHttp.PostJSON(apiUrl, map[string]string{"user_ticket": userTicket})
	if err != nil {
		return nil, kernelError.New(0, err)
	}

	var result userDetailResult
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, kernelError.New(0, err)
	}
	if result.ErrCode != 0 {
		return nil, kernelError.NewWithApiError(result.ApiError)
	}

	return &result.UserDetail, nil
}
