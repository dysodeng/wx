package open_platform

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/kernel"
	"github.com/dysodeng/wx/official"
	"github.com/dysodeng/wx/support/cache"
)

// OpenPlatform 微信开放平台
type OpenPlatform struct {
	config *config
	option *option
}

func NewOpenPlatform(appId, secret, token, aesKey string, opts ...Option) *OpenPlatform {
	cfg := &config{
		appId:  appId,
		secret: secret,
		token:  token,
		aesKey: aesKey,
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

	return &OpenPlatform{
		config: cfg,
		option: o,
	}
}

// Server 服务端
func (open *OpenPlatform) Server() *base.Server {
	server := base.NewServer(open)
	server.Push(&ComponentVerifyTicket{}, kernel.GuardEventComponentVerifyTicket)
	return server
}

// Official 授权到开放平台的公众号
// @param appId string 公众号appID
// @param authorizerRefreshToken string 公众号授权刷新token
func (open *OpenPlatform) Official(appId, authorizerRefreshToken string) *official.Official {
	return official.NewOfficialWithOpenPlatform(
		appId,
		authorizerRefreshToken,
		open,
		official.WithCache(open.option.cache),
		official.WithCacheKeyPrefix(open.option.cacheKeyPrefix),
	)
}
