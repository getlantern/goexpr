package redis

import (
	"fmt"
	"io"

	"github.com/getlantern/goexpr"
)

type hget struct {
	Key   goexpr.Expr
	Field goexpr.Expr
}

func HGet(key goexpr.Expr, field goexpr.Expr) goexpr.Expr {
	return &hget{
		Key:   key,
		Field: field,
	}
}

func (e *hget) Eval(params goexpr.Params) interface{} {
	redisClient := getRedisClient()

	_key := e.Key.Eval(params)
	if _key == nil {
		return nil
	}
	key := _key.(string)
	_field := e.Field.Eval(params)
	if _field == nil {
		return nil
	}
	field := fmt.Sprint(_field)

	// Check cache
	cacheMx.Lock()
	cache, cacheFound := caches[key]
	if !cacheFound {
		cache = newCache(key, cacheSize, func() (func(onUpdate func(key interface{}, value interface{})), error) {
			names, err := redisClient.HGetAll(key).Result()
			if err == io.EOF {
				return noopRefresh, nil
			} else if err != nil {
				return nil, err
			}
			return func(onUpdate func(key interface{}, value interface{})) {
				for k, v := range names {
					onUpdate(k, v)
				}
			}, nil
		})
		caches[key] = cache
	}
	cacheMx.Unlock()
	cached, cachedFound := cache.Get(field)
	if cachedFound {
		return cached
	}

	value, _ := redisClient.HGet(key, field).Result()
	cache.Add(field, value)
	return value
}

func (e *hget) WalkParams(cb func(string)) {
	e.Key.WalkParams(cb)
	e.Field.WalkParams(cb)
}

func (e *hget) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *hget) WalkLists(cb func(goexpr.List)) {
	e.Key.WalkLists(cb)
}

func (e *hget) String() string {
	return fmt.Sprintf("HGET(%v,%v)", e.Key, e.Field)
}
