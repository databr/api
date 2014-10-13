package database

import (
	"log"

	"github.com/databr/api/config"
	memcache "github.com/dustin/gomemcached/client"
)

func NewMemcache() *memcache.Client {
	m, err := memcache.Connect("tcp", config.MemcacheURL)
	if err != nil {
		panic(err)
	}

	if config.MemcacheUsername != "" {
		m.Auth(config.MemcacheUsername, config.MemcachePassword)
	}

	log.Println("Memcache Healthy", m.IsHealthy())
	log.Println(m.Set(0, "test", 0, 5, []byte("Testing...")))
	log.Println(m.Get(0, "test"))

	return m
}
