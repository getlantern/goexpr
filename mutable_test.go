package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMutable(t *testing.T) {
	e := Mutable()
	assert.Equal(t, "Mutable(<unset>)", e.String())
	assert.Nil(t, e.Eval(nil))
	e.Set(Constant(5))
	assert.Equal(t, "Mutable(5)", e.String())
	assert.Equal(t, 5, e.Eval(nil))
}
