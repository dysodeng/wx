package official

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/base/jssdk"
	"github.com/dysodeng/wx/official/article"
	"github.com/dysodeng/wx/official/oauth"
	"github.com/dysodeng/wx/official/template_message"
	"github.com/dysodeng/wx/official/user"
	"github.com/dysodeng/wx/support/cache"
)

// Official 公众号
type Official struct {
	config *config
	option *option
}

func NewOfficial(appId, appSecret, token, aesKey string, opts ...Option) (*Official, error) {
	c := &config{
		isOpenPlatform: false,
		appId:          appId,
		appSecret:      appSecret,
		token:          token,
		aesKey:         aesKey,
	}

	o := &option{
		cacheKeyPrefix: DefaultCacheKeyPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.cache == nil {
		o.cache = cache.NewMemoryCache()
	}

	return &Official{
		config: c,
		option: o,
	}, nil
}

// NewOfficialWithOpenPlatform 开放平台代公众号调用接口
func NewOfficialWithOpenPlatform(
	appId,
	authorizerRefreshToken string,
	authorizerAccount base.AuthorizerAccountInterface,
	opts ...Option,
) *Official {
	c := &config{
		isOpenPlatform:         true,
		appId:                  appId,
		authorizerRefreshToken: authorizerRefreshToken,
		authorizerAccount:      authorizerAccount,
	}

	o := &option{
		cacheKeyPrefix: DefaultCacheKeyPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}
	if o.cache == nil {
		o.cache = cache.NewMemoryCache()
	}

	return &Official{
		config: c,
		option: o,
	}
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
	return jssdk.NewJssdk(official)
}
