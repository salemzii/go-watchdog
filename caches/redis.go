package caches

import (
	"context"
	"os"
	"time"

	cache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

var mycache *cache.Cache

func MakeRedisCacheCheck(c *Cache) map[string]string {

	rdb := redis.NewClient(&redis.Options{
		//Addr: "localhost:6379",
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
	})

	mycache = cache.New(&cache.Options{
		Redis: rdb,
		// cache 1,000 keys for 1 minute
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return getItem()
}

func setItem() map[string]string {
	ctx := context.TODO()
	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   "watchdog",
		Value: "Item Set",
		TTL:   time.Minute * 5,
	}); err != nil {
		HandleCacheErr(err)
	}
	return map[string]string{"status": "ok"}
}

func getItem() map[string]string {
	ctx := context.TODO()
	if err := mycache.Get(ctx, "watchdog", map[string]string{}); err != nil {
		if err == cache.ErrCacheMiss {
			setItem()
		}
		HandleCacheErr(err)
	}
	return map[string]string{"status": "ok"}
}
