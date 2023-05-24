package open_platform

import (
	"github.com/dysodeng/wx/base/server"
	"github.com/dysodeng/wx/kernel/event"
	"github.com/dysodeng/wx/mini_program"
	"github.com/dysodeng/wx/official"
	"github.com/dysodeng/wx/open_platform/authorizer"
	"github.com/dysodeng/wx/open_platform/code"
	"github.com/dysodeng/wx/open_platform/code_template"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
)

// OpenPlatform 微信开放平台
type OpenPlatform struct {
	config *config
	option *option
}

// New 微信开放平台sdk实例
func New(appId, appSecret, token, aesKey string, opts ...Option) *OpenPlatform {
	cfg := &config{
		appId:     appId,
		appSecret: appSecret,
		token:     token,
		aesKey:    aesKey,
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

	return &OpenPlatform{
		config: cfg,
		option: o,
	}
}

// Server 服务端
func (open *OpenPlatform) Server() *server.Server {
	svr := server.New(open)
	svr.Register(&componentVerifyTicket{}, event.ComponentVerifyTicket)
	return svr
}

// Authorizer 公众账号授权
func (open *OpenPlatform) Authorizer() *authorizer.Authorizer {
	return authorizer.New(open)
}

// CodeTemplate 小程序代码模板
func (open *OpenPlatform) CodeTemplate() *code_template.CodeTemplate {
	return code_template.New(open)
}

// Code 小程序代码管理
func (open *OpenPlatform) Code() *code.Code {
	return code.New(open)
}

// Official 授权到开放平台的公众号
// @param appId string 公众号appID
// @param authorizerRefreshToken string 公众号授权刷新token
func (open *OpenPlatform) Official(appId, authorizerRefreshToken string) *official.Official {
	return official.NewWithOpenPlatform(
		appId,
		authorizerRefreshToken,
		open.config.token,
		open.config.aesKey,
		open,
		official.WithCache(open.option.cache),
		official.WithCacheKeyPrefix(open.option.cacheKeyPrefix),
		official.WithLocker(open.option.locker.Clone("official_"+appId)),
	)
}

// MiniProgram 授权到开放平台的小程序
// @param appId string 小程序appID
// @param authorizerRefreshToken string 小程序授权刷新token
func (open *OpenPlatform) MiniProgram(appId, authorizerRefreshToken string) *mini_program.MiniProgram {
	return mini_program.NewWithOpenPlatform(
		appId,
		authorizerRefreshToken,
		open.config.token,
		open.config.aesKey,
		open,
		mini_program.WithCache(open.option.cache),
		mini_program.WithCacheKeyPrefix(open.option.cacheKeyPrefix),
		mini_program.WithLocker(open.option.locker.Clone("mini_program_"+appId)),
	)
}
