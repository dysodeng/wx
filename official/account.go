package official

func (official *Official) AccountToken() string {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.AuthorizerAccountToken()
	}
	return official.config.token
}

func (official *Official) AccountAesKey() string {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.AuthorizerAccountAesKey()
	}
	return official.config.aesKey
}

func (official *Official) AccountAppId() string {
	return official.config.appId
}

func (official *Official) AccountAppSecret() string {
	return official.config.appSecret
}

func (official *Official) ComponentAppId() string {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.ComponentAppId()
	} else {
		return ""
	}
}

func (official *Official) ComponentAccessToken() string {
	if official.config.isOpenPlatform {
		return official.config.authorizerAccount.ComponentAccessToken()
	} else {
		return ""
	}
}

// IsOpenPlatform 是否为开放平台下的公众账号
func (official *Official) IsOpenPlatform() bool {
	return official.config.isOpenPlatform
}
