package lock

// Locker 锁接口，应用于单线程获取公众账号AccessToken等场景
type Locker interface {
	Lock() error
	Unlock() error
	Clone() Locker
}
