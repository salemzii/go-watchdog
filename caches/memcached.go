package caches

import "github.com/bradfitz/gomemcache/memcache"

func MakeMemcachedCheck(c *Cache) map[string]string {

	cache := memcache.New(c.Addrs)

	if err := cache.Ping(); err != nil {
		return HandleCacheErr("memcached", err)
	}
	return map[string]string{"status": "ok", "service": "memcached"}
}
