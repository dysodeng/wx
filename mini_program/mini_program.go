package mini_program

import (
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/mini_program/auth"
	"github.com/dysodeng/wx/mini_program/qr_code"
	"github.com/dysodeng/wx/mini_program/wxa_code"
	"github.com/dysodeng/wx/support/cache"
)

// MiniProgram 小程序
type MiniProgram struct {
	config *config
	option *option
}

func NewMiniProgram(appId, appSecret string, opts ...Option) *MiniProgram {
	c := &config{
		isOpenPlatform: false,
		appId:          appId,
		appSecret:      appSecret,
	}

	o := &option{
		cacheKeyPrefix: cache.DefaultCacheKeyPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.cache == nil {
		o.cache = cache.NewMemoryCache()
	}

	return &MiniProgram{
		config: c,
		option: o,
	}
}

// NewMiniProgramWithOpenPlatform 开放平台代小程序调用接口
func NewMiniProgramWithOpenPlatform(
	appId,
	authorizerRefreshToken string,
	authorizerAccount contracts.AuthorizerAccountInterface,
	opts ...Option,
) *MiniProgram {
	c := &config{
		isOpenPlatform:         true,
		appId:                  appId,
		authorizerRefreshToken: authorizerRefreshToken,
		authorizerAccount:      authorizerAccount,
	}

	o := &option{
		cacheKeyPrefix: cache.DefaultCacheKeyPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.cache == nil {
		o.cache = cache.NewMemoryCache()
	}

	return &MiniProgram{
		config: c,
		option: o,
	}
}

// Auth 用户登录
func (mp *MiniProgram) Auth() *auth.Auth {
	return auth.NewAuth(mp)
}

// WxaCode 小程序码
func (mp *MiniProgram) WxaCode() *wxa_code.WxaCode {
	return wxa_code.NewWxaCode(mp)
}

// QrCode 普通链接二维码
func (mp *MiniProgram) QrCode() *qr_code.QrCode {
	return qr_code.NewQrCode(mp)
}
