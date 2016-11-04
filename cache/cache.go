package cache

import "errors"

var (
	ErrCacheMiss  = errors.New("cache: key not found.")
	ErrNotStored  = errors.New("cache: not stored.")
	ErrNotSupport = errors.New("cache: not support.")
)

type CacheStore interface {
	GetString(key string) (string, error)
	GetInt(key string) (int, error)
	Set(key string, value interface{}) error
	Delete(key string) error
	Increment(key string, data uint64) (uint64, error)
	Decrement(key string, data uint64) (uint64, error)
	HashSet(key string, values []string) error
}
