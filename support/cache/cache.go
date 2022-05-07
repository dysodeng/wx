package cache

import "time"

const DefaultCacheKeyPrefix = "dy.wx.cache."

// Cache 缓存接口
type Cache interface {
	IsExist(key string) bool
	Get(key string) (interface{}, error)
	Put(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
	ClearAll() error
}
