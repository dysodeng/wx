package authorizer

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

// Info 授权账号详情信息
type Info struct {
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
		Idc           int            `json:"idc"`
		Signature     string         `json:"signature"`

		// 小程序独有字段
		RegisterType    int              `json:"register_type"`
		BasicConfig     map[string]bool  `json:"basic_config"`
		MiniProgramInfo *MiniProgramInfo `json:"MiniProgramInfo"`
	} `json:"authorizer_info"`
}

type MiniProgramInfo struct {
	Network     map[string][]string `json:"network"`
	Categories  []Categories        `json:"categories"`
	VisitStatus int                 `json:"visit_status"`
}

type Categories struct {
	First  string `json:"first"`
	Second string `json:"second"`
}
