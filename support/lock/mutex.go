package lock

import "sync"

// Mutex 通过 sync.Mutex 实现的锁
type Mutex struct {
	mu sync.Mutex
}

func (mu *Mutex) Lock() error {
	mu.mu.Lock()
	return nil
}

func (mu *Mutex) Unlock() error {
	mu.mu.Unlock()
	return nil
}

func (mu *Mutex) Clone() Locker {
	return &Mutex{}
}
