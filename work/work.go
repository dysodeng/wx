package work

import (
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
	"github.com/dysodeng/wx/work/base"
	"github.com/dysodeng/wx/work/mini_program"
)

// Work 企业微信
type Work struct {
	config *config
	option *option
}

func New(corpId, secret, token, aesKey string, opts ...Option) *Work {
	c := &config{
		corpId: corpId,
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
	if o.locker == nil {
		o.locker = &lock.Mutex{}
	}

	return &Work{
		config: c,
		option: o,
	}
}

// Base 基础模块
func (w *Work) Base() *base.Base {
	return base.New(w)
}

// MiniProgram 小程序
func (w *Work) MiniProgram() *mini_program.MiniProgram {
	return mini_program.New(w.config.corpId, w.config.secret, w.config.token, w.config.aesKey)
}
