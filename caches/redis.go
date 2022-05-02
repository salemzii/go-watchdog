package caches

import (
	"context"
	"time"

	cache "github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"github.com/salemzii/go-watchdog/service"
)

var mycache *cache.Cache

func MakeRedisCacheCheck(c *Cache) service.ServiceCheck {

	rdb := redis.NewClient(&redis.Options{
		//Addr: "localhost:6379",
		Addr:     c.Addrs,    //os.Getenv("REDIS_ADDR"),
		Password: c.Password, //os.Getenv("REDIS_PASSWORD"),
	})

	mycache = cache.New(&cache.Options{
		Redis: rdb,
		// cache 1,000 keys for 1 minute
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return ping(rdb)
}

func ping(rdb *redis.Client) service.ServiceCheck {
	ctx := context.TODO()

	status := rdb.Ping(ctx)
	if status.Err() != nil {
		return service.HandleError("redis", status.Err())
	}
	return service.HandleSuccess("redis", nil)
}

func setItem() map[string]string {
	ctx := context.TODO()
	if err := mycache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   "watchdog",
		Value: "Item Set",
		TTL:   time.Minute * 5,
	}); err != nil {
		HandleCacheErr("redis", err)
	}
	return map[string]string{"status": "ok"}
}

func getItem() map[string]string {
	ctx := context.TODO()
	if err := mycache.Get(ctx, "watchdog", map[string]string{}); err != nil {
		if err == cache.ErrCacheMiss {
			setItem()
		} else {
			HandleCacheErr("redis", err)
		}
	}
	return map[string]string{"status": "ok"}
}
