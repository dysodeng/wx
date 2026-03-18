package work

import (
	"github.com/dysodeng/wx/support/cache"
)

func (w *Work) Token() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.AuthorizerAccountToken()
	}
	return w.config.token
}

func (w *Work) AesKey() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.AuthorizerAccountAesKey()
	}
	return w.config.aesKey
}

func (w *Work) AppId() string {
	return w.config.corpId
}

func (w *Work) AppSecret() string {
	return w.config.secret
}

func (w *Work) ComponentAppId() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.ComponentAppId()
	}
	return ""
}

func (w *Work) ComponentAccessToken() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.ComponentAccessToken()
	}
	return ""
}

func (w *Work) IsOpenPlatform() bool {
	return false
}

func (w *Work) PlatformType() string {
	return "work"
}

func (w *Work) Cache() (cache.Cache, string) {
	return w.option.cache, w.option.cacheKeyPrefix
}
