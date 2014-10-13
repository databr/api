package database

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/databr/api/config"
)

func NewMemcache() *memcache.Client {
	m := memcache.New(config.MemcacheURL)
	m.Set(&memcache.Item{Key: "test", Value: []byte("tested")})
	_, err := m.Get("test")
	if err != nil && err != memcache.ErrCacheMiss {
		panic(err)
	}

	return m
}
