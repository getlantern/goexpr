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

	assert.Equal(t, "", Concat(pipe).Eval(nil))
	assert.Equal(t, "a", Concat(pipe, a).Eval(nil))
	assert.Equal(t, "a|b", Concat(pipe, a, b).Eval(nil))
	assert.Equal(t, "a|b|c", Concat(pipe, a, b, c).Eval(nil))

	assert.Equal(t, "CONCAT(|)", Concat(pipe).String())
	assert.Equal(t, "CONCAT(|, a)", Concat(pipe, a).String())
	assert.Equal(t, "CONCAT(|, a, b)", Concat(pipe, a, b).String())
	assert.Equal(t, "CONCAT(|, a, b, c)", Concat(pipe, a, b, c).String())
}
