package redis

import (
	"fmt"
	"github.com/getlantern/goexpr"
	"github.com/getlantern/golog"
	"github.com/hashicorp/golang-lru"
	"gopkg.in/redis.v3"
	"sync"
	"time"
)

var (
	log = golog.LoggerFor("goexpr.redis")

	redisClient *redis.Client
	caches      map[string]*lru.Cache
	cacheSize   int
	cacheMx     sync.RWMutex
)

func Configure(client *redis.Client, maxCacheSize int) {
	redisClient = client
	caches = make(map[string]*lru.Cache)
	cacheSize = maxCacheSize
	go warmCaches()
}

func warmCaches() {
	for {
		cacheMx.Lock()
		cachesCopy := make(map[string]*lru.Cache, len(caches))
		for hash, cache := range caches {
			cachesCopy[hash] = cache
		}
		cacheMx.Unlock()
		for hash, cache := range cachesCopy {
			names, err := redisClient.HGetAllMap(hash).Result()
			if err != nil {
				log.Errorf("Unable to load hash %v from Redis: %v", hash, err)
			} else {
				cacheMx.Lock()
				cache.Purge()
				for k, v := range names {
					cache.Add(k, v)
				}
				cacheMx.Unlock()
			}
		}
		time.Sleep(5 * time.Minute)
	}
}

func HMGet(hash goexpr.Expr, key goexpr.Expr) goexpr.Expr {
	return &hmget{
		hash: hash,
		key:  key,
	}
}

type hmget struct {
	hash goexpr.Expr
	key  goexpr.Expr
}

func (e *hmget) Eval(params goexpr.Params) interface{} {
	_hash := e.hash.Eval(params)
	if _hash == nil {
		return nil
	}
	hash := _hash.(string)
	key := e.key.Eval(params)
	if key == nil {
		return nil
	}

	// Check cache
	cacheMx.Lock()
	cache, cacheFound := caches[hash]
	if !cacheFound {
		cache, _ = lru.New(cacheSize)
		caches[hash] = cache
	}
	cacheMx.Unlock()
	cacheMx.RLock()
	cached, cachedFound := cache.Get(key)
	if cachedFound {
		cacheMx.RUnlock()
		return cached
	}
	cacheMx.RUnlock()

	keyString := fmt.Sprint(key)
	value, _ := redisClient.HGet(hash, keyString).Result()
	cache.Add(key, value)
	return value
}

func (e *hmget) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *hmget) WalkLists(cb func(goexpr.List)) {
	e.key.WalkLists(cb)
}

func (e *hmget) String() string {
	return fmt.Sprintf("HMGET(%v,%v)", e.hash, e.key)
}
