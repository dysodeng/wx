package open_platform

func (open *OpenPlatform) AccountToken() string {
	return open.config.token
}

func (open *OpenPlatform) AccountAesKey() string {
	return open.config.aesKey
}

func (open *OpenPlatform) AccountAppId() string {
	return open.config.appId
}

func (open *OpenPlatform) AccountAppSecret() string {
	return open.config.secret
}

func (open *OpenPlatform) ComponentAppId() string {
	return open.config.appId
}

func (open *OpenPlatform) ComponentAccessToken() string {
	token, err := open.AccessToken(false)
	if err != nil {
		return ""
	}
	return token.AccessToken
}

func (open *OpenPlatform) IsOpenPlatform() bool {
	return false
}

func (open *OpenPlatform) AuthorizerAccountToken() string {
	// TODO
	return ""
}

func (open *OpenPlatform) AuthorizerAccountAesKey() string {
	// TODO
	return ""
}
