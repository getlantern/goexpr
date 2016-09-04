package goexpr

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestLazy(t *testing.T) {
	e := Lazy(func() interface{} {
		return rand.Int()
	})
	val := e.Eval(nil)
	assert.Equal(t, val, e.Eval(nil), "Lazy expression should always return the first value")
	assert.Equal(t, "LAZY(", e.String()[:5])
}
