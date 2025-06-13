package cache

import "time"

type CacheContext interface {
	// Get retrieves the value associated with the given key.
	Get(key string) (string, error)

	// Set sets the value for the given key.
	Set(key string, value string) error

	// SetWithTTL sets the value with a time-to-live duration.
	SetWithTTL(key string, value string, ttl time.Duration) error

	// Delete removes the value associated with the key.
	Delete(key string) error

	// Exists checks if the key exists in the cache.
	Exists(key string) (bool, error)

	// Incr atomically increments the integer value of a key by 1.
	Incr(key string) (int64, error)

	// Decr atomically decrements the integer value of a key by 1.
	Decr(key string) (int64, error)
}
