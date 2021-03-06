package caches

import (
	"log"
	"strings"

	"github.com/salemzii/go-watchdog/service"
)

var supportedCaches = map[string][]string{
	"cache": {"redis", "memcache", "couchbase", ""},
}

//https://developer20.com/garnish-simple-varnish-in-go/

type Cache struct {
	Type     string `json:"type"`
	Addrs    string `json:"addrs"`
	UriOnly  string `json:"uri_only"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cache *Cache) Uri_Only() bool {
	return cache.UriOnly != ""
}

func (cache *Cache) GetCacheDriver() service.ServiceCheck {

	switch strings.ToLower(cache.Type) {

	case "redis":
		return MakeRedisCacheCheck(cache)
	case "memcached":
		return MakeMemcachedCheck(cache)
	case "couchbase":
		// get couchbase-driver
	case "varnish":
		// get varnish-driver
	default:
		log.Println("Cache " + cache.Type + " not supported")
	}
	return service.ServiceCheck{}
}

func HandleCacheErr(service string, err error) map[string]string {
	status := map[string]string{
		"status":  "Failed",
		"error":   err.Error(),
		"service": service,
	}
	return status
}
