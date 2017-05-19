package redis

import (
	"sync"

	"github.com/getlantern/golog"
	"gopkg.in/redis.v5"
	"gopkg.in/vmihailenco/msgpack.v2"
)

var (
	log = golog.LoggerFor("goexpr.redis")

	redisClient *redis.Client
	caches      map[string]*cache
	cacheSize   int
	cacheMx     sync.Mutex
)

func init() {
	msgpack.RegisterExt(90, &hget{})
}

func Configure(client *redis.Client, maxCacheSize int) {
	redisClient = client
	caches = make(map[string]*cache)
	cacheSize = maxCacheSize
}
