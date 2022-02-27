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
