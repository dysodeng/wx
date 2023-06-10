package app

import (
	"github.com/dysodeng/wx/app/oauth"
	"github.com/dysodeng/wx/support/cache"
	"github.com/dysodeng/wx/support/lock"
)

// App 移动应用
type App struct {
	config *config
	option *option
}

func New(appId, appSecret, token, aesKey string, opts ...Option) *App {
	c := &config{
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

	return &App{
		config: c,
		option: o,
	}
}

func (app *App) OAuth() *oauth.OAuth {
	return oauth.New(app, oauth.WithLocker(app.option.locker.Clone("wx_app_oauth")))
}
