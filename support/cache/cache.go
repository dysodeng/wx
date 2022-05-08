package cache

import "time"

const DefaultCacheKeyPrefix = "dy.wx.cache."

// Cache 缓存接口
type Cache interface {
	IsExist(key string) bool
	Get(key string) (string, error)
	Put(key string, value string, expiration time.Duration) error
	Delete(key string) error
	ClearAll() error
}
