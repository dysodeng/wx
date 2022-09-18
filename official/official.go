package official

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/base/jssdk"
	"github.com/dysodeng/wx/kernel/contracts"
	"github.com/dysodeng/wx/official/article"
	"github.com/dysodeng/wx/official/authorizer"
	"github.com/dysodeng/wx/official/oauth"
	"github.com/dysodeng/wx/official/qr_code"
	"github.com/dysodeng/wx/official/template_message"
	"github.com/dysodeng/wx/official/user"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
)

// Official 公众号
type Official struct {
	config *config
	option *option
}

func NewOfficial(appId, appSecret, token, aesKey string, opts ...Option) *Official {
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

	return &Official{
		config: c,
		option: o,
	}
}

// NewOfficialWithOpenPlatform 开放平台代公众号调用接口
func NewOfficialWithOpenPlatform(
	appId,
	authorizerRefreshToken,
	token,
	aesKey string,
	authorizerAccount contracts.AuthorizerInterface,
	opts ...Option,
) *Official {
	c := &config{
		isOpenPlatform:         true,
		appId:                  appId,
		authorizerRefreshToken: authorizerRefreshToken,
		token:                  token,
		aesKey:                 aesKey,
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

	return &Official{
		config: c,
		option: o,
	}
}

// Authorizer 公众号授权开放平台后的相关接口
// 此类接口只能由授权到开放平台的公众号调用或直接由开放平台调用
func (official *Official) Authorizer() *authorizer.Authorizer {
	return authorizer.NewAuthorizer(official)
}

// Server 服务端
func (official *Official) Server() *base.Server {
	return base.NewServer(official)
}

// OAuth 用户授权
func (official *Official) OAuth() *oauth.OAuth {
	return oauth.NewOAuth(official)
}

// User 用户管理
func (official *Official) User() *user.User {
	return user.NewUser(official)
}

// UserTag 用户标签管理
func (official *Official) UserTag() *user.Tag {
	return user.NewUserTag(official)
}

// Article 文章管理
func (official *Official) Article() *article.Article {
	return article.NewArticle(official)
}

// TemplateMessage 模板消息
func (official *Official) TemplateMessage() *template_message.TemplateMessage {
	return template_message.NewTemplateMessage(official)
}

// Jssdk 微信JSSDK
func (official *Official) Jssdk() *jssdk.Jssdk {
	return jssdk.NewJssdk(official, jssdk.WithLocker(official.option.locker.Clone("jssdk")))
}

// QrCode 带参数的二维码
func (official *Official) QrCode() *qr_code.QrCode {
	return qr_code.NewQrCode(official)
}
