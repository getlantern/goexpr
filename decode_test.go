package goexpr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode(t *testing.T) {
	e := Decode(Constant(1))
	assert.Nil(t, e.Eval(nil))
	assert.Equal(t, "DECODE(1)", e.String())

	e = Decode(Constant(1), Constant(5))
	assert.Equal(t, 5, e.Eval(nil))
	assert.Equal(t, "DECODE(1, 5)", e.String())

	e = Decode(Constant(1), Constant(2), Constant(2), Constant(3), Constant(3), Constant(5))
	assert.Equal(t, 5, e.Eval(nil))
	assert.Equal(t, "DECODE(1, 2, 2, 3, 3, 5)", e.String())

	e = Decode(Constant(1), Constant(2), Constant(2), Constant(3), Constant(3), Constant(1), Constant(1), Constant(5))
	assert.Equal(t, 1, e.Eval(nil))
	assert.Equal(t, "DECODE(1, 2, 2, 3, 3, 1, 1, 5)", e.String())
}
