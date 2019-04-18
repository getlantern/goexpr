package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplit(t *testing.T) {
	pipe := Constant("|")
	dot := Constant(".")
	source := Constant("a|b|c")

	assert.Equal(t, "a", msgpacked(t, Split(source, pipe, Constant(0))).Eval(nil))
	assert.Equal(t, "b", msgpacked(t, Split(source, pipe, Constant(1))).Eval(nil))
	assert.Equal(t, "c", msgpacked(t, Split(source, pipe, Constant(2))).Eval(nil))
	assert.Equal(t, "c", msgpacked(t, Split(source, pipe, Constant(-1))).Eval(nil))
	assert.Equal(t, "b", msgpacked(t, Split(source, pipe, Constant(-2))).Eval(nil))
	assert.Equal(t, "a", msgpacked(t, Split(source, pipe, Constant(-3))).Eval(nil))
	assert.Nil(t, msgpacked(t, Split(source, pipe, Constant(3))).Eval(nil))
	assert.Nil(t, msgpacked(t, Split(source, pipe, Constant(-4))).Eval(nil))
	assert.Equal(t, "a|b|c", msgpacked(t, Split(source, dot, Constant(0))).Eval(nil))
}
