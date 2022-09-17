package cache

import "github.com/pkg/errors"

var (
	ErrKeyExpired  = errors.New("key expired")
	ErrKeyNotExist = errors.New("key not exist")
)
