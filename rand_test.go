package goexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRand(t *testing.T) {
	rand := Rand(Param("a"))
	params := MapParams{"a": 0.5} // 50%
	countOfTrue := float64(0)
	iters := 1000000
	for i := 0; i < iters; i++ {
		if rand.Eval(params).(bool) {
			countOfTrue++
		}
	}
	ratioOfTrue := countOfTrue / float64(iters)
	assert.True(t, ratioOfTrue > 0.4)
	assert.True(t, ratioOfTrue < 0.6)
}
