package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	pipe := Constant("|")
	a := Constant("a")
	b := Constant("b")
	c := Constant("c")

	assert.Equal(t, "", msgpacked(t, Concat(pipe)).Eval(nil))
	assert.Equal(t, "a", msgpacked(t, Concat(pipe, a)).Eval(nil))
	assert.Equal(t, "a|b", msgpacked(t, Concat(pipe, a, b)).Eval(nil))
	assert.Equal(t, "a|b|c", msgpacked(t, Concat(pipe, a, b, c)).Eval(nil))

	assert.Equal(t, "CONCAT(|)", msgpacked(t, Concat(pipe)).String())
	assert.Equal(t, "CONCAT(|, a)", msgpacked(t, Concat(pipe, a)).String())
	assert.Equal(t, "CONCAT(|, a, b)", msgpacked(t, Concat(pipe, a, b)).String())
	assert.Equal(t, "CONCAT(|, a, b, c)", msgpacked(t, Concat(pipe, a, b, c)).String())
}
