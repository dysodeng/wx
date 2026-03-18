package mini_program

import (
	"github.com/dysodeng/wx/support/cache"
)

func (w *MiniProgram) Token() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.AuthorizerAccountToken()
	}
	return w.config.token
}

func (w *MiniProgram) AesKey() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.AuthorizerAccountAesKey()
	}
	return w.config.aesKey
}

func (w *MiniProgram) AppId() string {
	return w.config.corpId
}

func (w *MiniProgram) AppSecret() string {
	return w.config.secret
}

func (w *MiniProgram) ComponentAppId() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.ComponentAppId()
	}
	return ""
}

func (w *MiniProgram) ComponentAccessToken() string {
	if w.IsOpenPlatform() {
		return w.config.authorizerAccount.ComponentAccessToken()
	}
	return ""
}

func (w *MiniProgram) IsOpenPlatform() bool {
	return false
}

func (w *MiniProgram) PlatformType() string {
	return "mini_program"
}

func (w *MiniProgram) Cache() (cache.Cache, string) {
	return w.option.cache, w.option.cacheKeyPrefix
}
