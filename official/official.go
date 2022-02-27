package official

import (
	"github.com/dysodeng/wx/base"
	"github.com/dysodeng/wx/base/cache"
	"github.com/dysodeng/wx/official/article"
	"github.com/dysodeng/wx/official/user"
)

// Official 公众号
type Official struct {
	config *config
	option *option
}

func NewOfficial(cfg Config, opts ...Option) (*Official, error) {
	c := &config{}
	cfg(c)

	o := &option{
		cache:          cache.NewMemoryCache(),
		cacheKeyPrefix: DefaultCacheKeyPrefix,
	}
	for _, opt := range opts {
		opt(o)
	}

	return &Official{
		config: c,
		option: o,
	}, nil
}

// Server 服务端
func (official *Official) Server() *base.Server {
	return base.NewServer(official)
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
