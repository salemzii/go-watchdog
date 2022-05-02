package caches

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/salemzii/go-watchdog/service"
)

func MakeMemcachedCheck(c *Cache) service.ServiceCheck {

	cache := memcache.New(c.Addrs)

	if err := cache.Ping(); err != nil {
		return service.HandleError("memcached", err)
	}
	return service.HandleSuccess("memcached", nil)
}
