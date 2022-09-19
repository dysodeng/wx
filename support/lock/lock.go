package lock

// Locker 锁接口，应用于并发获取公众账号AccessToken等场景
type Locker interface {
	Lock()
	Unlock()
	Clone(name string) Locker
}
