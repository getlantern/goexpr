package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubstr(t *testing.T) {
	source := Constant("abc")

	apply := func(from int, to int) interface{} {
		return Substr(source, Constant(from), Constant(to)).Eval(nil)
	}

	assert.Equal(t, "a", apply(0, 1))
	assert.Equal(t, "ab", apply(0, 2))
	assert.Equal(t, "abc", apply(0, 3))
	assert.Equal(t, "abc", apply(0, 0))

	assert.Equal(t, "b", apply(1, 1))
	assert.Equal(t, "bc", apply(1, 2))
	assert.Equal(t, "bc", apply(1, 0))

	assert.Equal(t, "c", apply(2, 1))
	assert.Equal(t, "c", apply(2, 0))

	assert.Nil(t, apply(3, 0))
}
