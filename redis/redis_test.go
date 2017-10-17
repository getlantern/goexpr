package redis

import (
	"testing"
	"time"

	"github.com/getlantern/goexpr"
	"github.com/getlantern/testredis"
	"github.com/stretchr/testify/assert"
)

const (
	testCacheInvalidationPeriod = 500 * time.Millisecond
)

func TestHGet(t *testing.T) {
	SetCacheInvalidationPeriod(testCacheInvalidationPeriod)
	defer SetCacheInvalidationPeriod(defaultCacheInvalidationPeriod)

	redis, err := testredis.Open()
	if !assert.NoError(t, err) {
		return
	}
	defer redis.Close()

	client := redis.Client()
	Configure(redis.Client(), 100)
	client.HSet("h1", "ka", 1)
	client.HSet("h1", "kb", 2)
	assert.Equal(t, "1", HGet(goexpr.Constant("h1"), goexpr.Constant("ka")).Eval(nil), "Should get initial ka")
	assert.Equal(t, "2", HGet(goexpr.Constant("h1"), goexpr.Constant("kb")).Eval(nil), "Should get initial kb")

	client.HSet("h1", "ka", 11)
	client.HSet("h1", "kb", 12)
	client.HSet("h1", "kc", 13)
	assert.Equal(t, "1", HGet(goexpr.Constant("h1"), goexpr.Constant("ka")).Eval(nil), "Should get cached ka")
	assert.Equal(t, "2", HGet(goexpr.Constant("h1"), goexpr.Constant("kb")).Eval(nil), "Should get cached kb")
	assert.Equal(t, "13", HGet(goexpr.Constant("h1"), goexpr.Constant("kc")).Eval(nil), "Should get newly added kc")

	time.Sleep(testCacheInvalidationPeriod * 2)
	assert.Equal(t, "11", HGet(goexpr.Constant("h1"), goexpr.Constant("ka")).Eval(nil), "Should get new ka")
	assert.Equal(t, "12", HGet(goexpr.Constant("h1"), goexpr.Constant("kb")).Eval(nil), "Should get new kb")
	assert.Equal(t, "13", HGet(goexpr.Constant("h1"), goexpr.Constant("kc")).Eval(nil), "Should still get newly added kc")
}

func TestSIsMember(t *testing.T) {
	SetCacheInvalidationPeriod(testCacheInvalidationPeriod)
	defer SetCacheInvalidationPeriod(defaultCacheInvalidationPeriod)

	redis, err := testredis.Open()
	if !assert.NoError(t, err) {
		return
	}
	defer redis.Close()

	client := redis.Client()
	Configure(redis.Client(), 100)
	client.SAdd("s1", "ka")
	assert.True(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("ka")).Eval(nil).(bool), "ka should be in initial set")
	assert.False(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("kb")).Eval(nil).(bool), "kb should not be in initial set")

	client.SAdd("s1", "kb")
	assert.True(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("ka")).Eval(nil).(bool), "ka should be in cached set")
	assert.True(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("kb")).Eval(nil).(bool), "kb should be in cached set")

	time.Sleep(testCacheInvalidationPeriod * 2)
	assert.True(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("ka")).Eval(nil).(bool), "ka should be in new set")
	assert.True(t, SIsMember(goexpr.Constant("s1"), goexpr.Constant("kb")).Eval(nil).(bool), "kb should be in new set")
}

func TestLua(t *testing.T) {
	script := goexpr.Constant(`return redis.call('SISMEMBER', KEYS[1], ARGV[1])`)
	SetCacheInvalidationPeriod(testCacheInvalidationPeriod)
	defer SetCacheInvalidationPeriod(defaultCacheInvalidationPeriod)

	redis, err := testredis.Open()
	if !assert.NoError(t, err) {
		return
	}
	defer redis.Close()

	client := redis.Client()
	Configure(redis.Client(), 100)
	client.SAdd("s1", "ka")
	assert.EqualValues(t, 1, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("ka")).Eval(nil).(int64), "ka should be in initial set")
	assert.EqualValues(t, 0, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("kb")).Eval(nil).(int64), "kb should not be in initial set")

	client.SAdd("s1", "kb")
	assert.EqualValues(t, 1, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("ka")).Eval(nil).(int64), "ka should be in cached set")
	assert.EqualValues(t, 0, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("kb")).Eval(nil).(int64), "kb should not be in cached set")

	time.Sleep(testCacheInvalidationPeriod * 2)
	assert.EqualValues(t, 1, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("ka")).Eval(nil).(int64), "ka should be in new set")
	assert.EqualValues(t, 1, Lua(script, []goexpr.Expr{goexpr.Constant("s1")}, goexpr.Constant("kb")).Eval(nil).(int64), "kb should be in new set")
}

func TestLuaStringFind(t *testing.T) {
	script := goexpr.Constant(`return string.find(redis.call('HGET', KEYS[1], ARGV[1]), ARGV[2])`)
	SetCacheInvalidationPeriod(testCacheInvalidationPeriod)
	defer SetCacheInvalidationPeriod(defaultCacheInvalidationPeriod)

	redis, err := testredis.Open()
	if !assert.NoError(t, err) {
		return
	}
	defer redis.Close()

	client := redis.Client()
	Configure(redis.Client(), 100)
	client.HSet("h1", "ka", "iamastring")

	assert.True(t, Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("ka"), goexpr.Constant("amas")).Eval(nil).(int64) > 0, "amas should be found")
	assert.Nil(t, Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("ka"), goexpr.Constant(`"lamas":`)).Eval(nil), "lamas should not be found")
}
