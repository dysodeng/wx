package mini_program

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/mini_program/auth"
	"github.com/dysodeng/wx/mini_program/authorizer"
	"github.com/dysodeng/wx/mini_program/content"
	"github.com/dysodeng/wx/mini_program/encryptor"
	"github.com/dysodeng/wx/mini_program/qr_code"
	"github.com/dysodeng/wx/mini_program/wxa_code"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
)

// MiniProgram 小程序
type MiniProgram struct {
	config *config
	option *option
}

func NewMiniProgram(appId, appSecret, token, aesKey string, opts ...Option) *MiniProgram {
	c := &config{
		isOpenPlatform: false,
		appId:          appId,
		appSecret:      appSecret,
		token:          token,
		aesKey:         aesKey,
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
	if o.locker == nil {
		o.locker = &lock.Mutex{}
	}

	return &MiniProgram{
		config: c,
		option: o,
	}
}

// NewMiniProgramWithOpenPlatform 开放平台代小程序调用接口
func NewMiniProgramWithOpenPlatform(
	appId,
	authorizerRefreshToken,
	token,
	aesKey string,
	authorizerAccount contracts.AuthorizerInterface,
	opts ...Option,
) *MiniProgram {
	c := &config{
		isOpenPlatform:         true,
		appId:                  appId,
		token:                  token,
		aesKey:                 aesKey,
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
	if o.locker == nil {
		o.locker = &lock.Mutex{}
	}

	return &MiniProgram{
		config: c,
		option: o,
	}
}

// Authorizer 小程序授权开放平台后的相关接口
// 此类接口只能由授权到开放平台的小程序调用或直接由开放平台调用
func (mp *MiniProgram) Authorizer() *authorizer.Authorizer {
	return authorizer.NewAuthorizer(mp)
}

// Server 服务端
func (mp *MiniProgram) Server() *base.Server {
	return base.NewServer(mp)
}

// Encryptor 小程序加密数据的解密
func (mp *MiniProgram) Encryptor() *encryptor.Encryptor {
	return encryptor.NewEncryptor(mp)
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

// ContentSecurity 内容安全
func (mp *MiniProgram) ContentSecurity() *content.Security {
	return content.NewContentSecurity(mp)
}
