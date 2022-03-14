package cache

import (
	"sync"
	"time"
)

// Item 缓存项
type item struct {
	value    interface{}
	created  time.Time
	lifetime time.Duration
}

// isExpire 缓存项是否过期
func (item *item) isExpire() bool {
	if item.lifetime == 0 {
		return false
	}

	return time.Since(item.created) > item.lifetime
}

const defaultGcTime = 60 // 1分钟gc一次

// MemoryCache 内存缓存
type MemoryCache struct {
	sync.RWMutex
	duration time.Duration
	items    map[string]*item
	gcTime   int
}

func NewMemoryCache() *MemoryCache {
	m := &MemoryCache{items: make(map[string]*item), gcTime: defaultGcTime, duration: time.Second * defaultGcTime}
	go m.gc()
	return m
}

// IsExist 缓存项是否存在
func (c *MemoryCache) IsExist(key string) bool {
	c.RLock()
	defer c.RUnlock()

	if i, ok := c.items[key]; ok {
		return !i.isExpire()
	}

	return false
}

// Get 获取缓存项
func (c *MemoryCache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	if i, ok := c.items[key]; ok {
		if !i.isExpire() {
			return nil, ErrKeyExpired
		}

		return i.value, nil
	}

	return nil, ErrKeyNotExist
}

// Put 设置缓存项
func (c *MemoryCache) Put(key string, value interface{}, expiration time.Duration) error {
	c.Lock()
	defer c.Unlock()

	c.items[key] = &item{
		value:    value,
		created:  time.Now(),
		lifetime: expiration,
	}

	return nil
}

// Delete 删除删除项
func (c *MemoryCache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	delete(c.items, key)

	return nil
}

// ClearAll 清除所有缓存项
func (c *MemoryCache) ClearAll() error {
	c.Lock()
	defer c.Unlock()

	c.items = make(map[string]*item)

	return nil
}

// gc 定时清除已过期缓存
func (c *MemoryCache) gc() {
	c.RLock()
	gcTime := c.gcTime
	c.RUnlock()

	if gcTime < 1 {
		return
	}

	for {
		<-time.After(c.duration)
		c.RLock()
		if c.items == nil {
			c.RUnlock()
			return
		}
		c.RUnlock()

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)
		}
	}
}

func (c *MemoryCache) expiredKeys() (keys []string) {
	c.RLock()
	defer c.RUnlock()
	for key, i := range c.items {
		if i.isExpire() {
			keys = append(keys, key)
		}
	}
	return
}

func (c *MemoryCache) clearItems(keys []string) {
	c.Lock()
	defer c.Unlock()
	for _, key := range keys {
		delete(c.items, key)
	}
}
