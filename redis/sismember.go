package redis

import (
	"fmt"
	"io"

	"github.com/getlantern/goexpr"
)

type sismember struct {
	Key    goexpr.Expr
	Member goexpr.Expr
}

func SIsMember(key goexpr.Expr, member goexpr.Expr) goexpr.Expr {
	return &sismember{
		Key:    key,
		Member: member,
	}
}

func (e *sismember) Eval(params goexpr.Params) interface{} {
	redisClient := getRedisClient()

	_key := e.Key.Eval(params)
	if _key == nil {
		return nil
	}
	key := _key.(string)
	_member := e.Member.Eval(params)
	if _member == nil {
		return nil
	}
	member := fmt.Sprint(_member)

	// Check cache
	cacheMx.Lock()
	cache, cacheFound := caches[key]
	if !cacheFound {
		cache = newCache(key, cacheSize, func() (func(onUpdate func(key interface{}, value interface{})), error) {
			members, err := redisClient.SMembers(key).Result()
			if err == io.EOF {
				return noopRefresh, nil
			} else if err != nil {
				return nil, err
			}
			return func(onUpdate func(key interface{}, value interface{})) {
				for _, member := range members {
					onUpdate(member, true)
				}
			}, nil
		})
		caches[key] = cache
	}
	cacheMx.Unlock()
	cached, cachedFound := cache.Get(member)
	if cachedFound {
		return cached
	}

	value, _ := redisClient.SIsMember(key, member).Result()
	cache.Add(member, value)
	return value
}

func (e *sismember) WalkParams(cb func(string)) {
	e.Key.WalkParams(cb)
	e.Member.WalkParams(cb)
}

func (e *sismember) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *sismember) WalkLists(cb func(goexpr.List)) {
	e.Key.WalkLists(cb)
}

func (e *sismember) String() string {
	return fmt.Sprintf("SISMEMBER(%v,%v)", e.Key, e.Member)
}
