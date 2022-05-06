package open_platform

func (open *OpenPlatform) AccountToken() string {
	// TODO
	return ""
}

func (open *OpenPlatform) AccountAesKey() string {
	// TODO
	return ""
}

func (open *OpenPlatform) AccountAppId() string {
	// TODO
	return ""
}

func (open *OpenPlatform) AccountAppSecret() string {
	return open.config.appSecret
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
