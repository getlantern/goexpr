package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcat(t *testing.T) {
	pipe := Constant("|")
	a := Constant("a")
	b := Constant(nil)
	c := Constant("c")

	assert.Equal(t, "", msgpacked(t, Concat(pipe)).Eval(nil))
	assert.Equal(t, "a", msgpacked(t, Concat(pipe, a)).Eval(nil))
	assert.Equal(t, "a|", msgpacked(t, Concat(pipe, a, b)).Eval(nil))
	assert.Equal(t, "a||c", msgpacked(t, Concat(pipe, a, b, c)).Eval(nil))

	assert.Equal(t, "CONCAT(|)", msgpacked(t, Concat(pipe)).String())
	assert.Equal(t, "CONCAT(|, a)", msgpacked(t, Concat(pipe, a)).String())
	assert.Equal(t, "CONCAT(|, a, <nil>)", msgpacked(t, Concat(pipe, a, b)).String())
	assert.Equal(t, "CONCAT(|, a, <nil>, c)", msgpacked(t, Concat(pipe, a, b, c)).String())
}

func TestWalkOneToOneConcat(t *testing.T) {
	e := msgpacked(t, Concat(Param("a"), Param("b")))

	found := make(map[string]bool)
	e.WalkOneToOneParams(func(name string) {
		found[name] = true
	})
	assert.Empty(t, found)

	pe := msgpacked(t, P(e))
	found = make(map[string]bool)
	pe.WalkOneToOneParams(func(name string) {
		found[name] = true
	})
	assert.True(t, found["a"])
	assert.True(t, found["b"])
	assert.Equal(t, "P(CONCAT(a, b))", pe.String())
}
