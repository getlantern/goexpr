package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	pipe := Constant("|")
	source := Constant("a|b|c")

	assert.Equal(t, "a", Split(source, pipe, Constant(0)).Eval(nil))
	assert.Equal(t, "b", Split(source, pipe, Constant(1)).Eval(nil))
	assert.Equal(t, "c", Split(source, pipe, Constant(2)).Eval(nil))
	assert.Nil(t, Split(source, pipe, Constant(3)).Eval(nil))
}
