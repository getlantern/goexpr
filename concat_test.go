package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	doTestConcat(t, Concat, "CONCAT")
}

func TestPConcat(t *testing.T) {
	doTestConcat(t, PConcat, "PCONCAT")
}

func doTestConcat(t *testing.T, fn func(...Expr) Expr, name string) {
	pipe := Constant("|")
	a := Constant("a")
	b := Constant(nil)
	c := Constant("c")

	assert.Equal(t, "", msgpacked(t, fn(pipe)).Eval(nil))
	assert.Equal(t, "a", msgpacked(t, fn(pipe, a)).Eval(nil))
	assert.Equal(t, "a|", msgpacked(t, fn(pipe, a, b)).Eval(nil))
	assert.Equal(t, "a||c", msgpacked(t, fn(pipe, a, b, c)).Eval(nil))

	assert.Equal(t, name+"(|)", msgpacked(t, fn(pipe)).String())
	assert.Equal(t, name+"(|, a)", msgpacked(t, fn(pipe, a)).String())
	assert.Equal(t, name+"(|, a, <nil>)", msgpacked(t, fn(pipe, a, b)).String())
	assert.Equal(t, name+"(|, a, <nil>, c)", msgpacked(t, fn(pipe, a, b, c)).String())
}
