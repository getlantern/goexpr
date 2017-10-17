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

func TestLuaComplex(t *testing.T) {
	script := goexpr.Constant(`
		local access_data = redis.call('HGET', KEYS[1], ARGV[1])
		if not access_data then
			return "unknown"
		end

		local kcp = string.find(access_data, "kcpsettings:")
		if kcp then
			return "kcp"
		end

		local pt = string.match(access_data, "pluggabletransport: '(%a+)'")
		if pt then
			return pt
		end

		return "https"`)
	SetCacheInvalidationPeriod(testCacheInvalidationPeriod)
	defer SetCacheInvalidationPeriod(defaultCacheInvalidationPeriod)

	redis, err := testredis.Open()
	if !assert.NoError(t, err) {
		return
	}
	defer redis.Close()

	client := redis.Client()
	Configure(redis.Client(), 100)
	client.HSet("h1", "https", `{
	pluggabletransport: ''
}`)
	client.HSet("h1", "lampshade", `{
	pluggabletransport: 'lampshade'
}`)
	client.HSet("h1", "kcp", `{
	pluggabletransport: ''
  kcpsettings: {"blah": 1}
}`)

	assert.Equal(t, "https", Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("https")).Eval(nil).(string), "should find https")
	assert.Equal(t, "lampshade", Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("lampshade")).Eval(nil).(string), "should find lampshade")
	assert.Equal(t, "kcp", Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("kcp")).Eval(nil).(string), "should find kcp")
	assert.Equal(t, "unknown", Lua(script, []goexpr.Expr{goexpr.Constant("h1")}, goexpr.Constant("blahblah")).Eval(nil).(string), "should find unknown")
}
