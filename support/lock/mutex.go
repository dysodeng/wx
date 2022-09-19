package lock

import "sync"

// Mutex 通过 sync.Mutex 实现的锁
type Mutex struct {
	mu sync.Mutex
}

func (mu *Mutex) Lock() {
	mu.mu.Lock()
}

func (mu *Mutex) Unlock() {
	mu.mu.Unlock()
}

func (mu *Mutex) Clone(name string) Locker {
	return &Mutex{}
}
