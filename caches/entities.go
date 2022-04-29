package caches

import (
	"log"
	"strings"
)

var supportedCaches = map[string][]string{
	"cache": {"redis", "memcache", "couchbase", ""},
}

type Cache struct {
	Type     string `json:"type"`
	UriOnly  string `json:"uri_only"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (cache *Cache) Uri_Only() bool {
	return cache.UriOnly != ""
}

func (cache *Cache) GetCacheDriver() {

	switch strings.ToLower(cache.Type) {

	case "redis":
		// Get redis-driver
	case "memcached":
	// get memcached-driver
	case "couchbase":
		// get couchbase-driver
	case "varnish":
		// get varnish-driver
	default:
		log.Println("Cache " + cache.Type + " not supported")
	}
}

func HandleCacheErr(err error) map[string]string {
	status := map[string]string{
		"status": "Failed",
		"error":  err.Error(),
	}
	return status
}
