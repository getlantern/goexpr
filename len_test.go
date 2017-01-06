package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	assert.Equal(t, 0, Len(Constant("")).Eval(nil))
	assert.Equal(t, 3, Len(Constant("abc")).Eval(nil))
	assert.Equal(t, 1, Len(Constant(0)).Eval(nil))
	assert.Nil(t, Len(Constant(nil)).Eval(nil))
}
