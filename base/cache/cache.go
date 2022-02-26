package cache

import "time"

type Cache interface {
	IsExist(key string) bool
	Get(key string) (interface{}, error)
	Put(key string, value interface{}, expiration time.Duration) error
	Delete(key string) error
	ClearAll() error
}
