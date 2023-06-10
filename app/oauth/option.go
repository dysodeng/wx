package oauth

import "github.com/dysodeng/wx/support/lock"

type Option func(auth *OAuth)

// WithLocker 设置锁
func WithLocker(locker lock.Locker) Option {
	return func(o *OAuth) {
		o.locker = locker
	}
}
