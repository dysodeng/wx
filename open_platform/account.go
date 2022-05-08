package open_platform

func (open *OpenPlatform) Token() string {
	return open.config.token
}

func (open *OpenPlatform) AesKey() string {
	return open.config.aesKey
}

func (open *OpenPlatform) AppId() string {
	return open.config.appId
}

func (open *OpenPlatform) AppSecret() string {
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
	return open.Token()
}

func (open *OpenPlatform) AuthorizerAccountAesKey() string {
	return open.AesKey()
}
