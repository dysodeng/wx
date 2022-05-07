package cache

import "errors"

var (
	ErrKeyExpired  = errors.New("key expired")
	ErrKeyNotExist = errors.New("key not exist")
)
