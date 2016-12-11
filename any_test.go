package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	e1 := Any(Constant(nil), Constant(""), Constant(1))
	e2 := Any(Constant(2), Constant(nil), Constant(""))
	assert.Equal(t, 1, e1.Eval(nil))
	assert.Equal(t, 2, e2.Eval(nil))
}
