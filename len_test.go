package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLen(t *testing.T) {
	assert.Equal(t, 0, msgpacked(t, Len(Constant(""))).Eval(nil))
	assert.Equal(t, 3, msgpacked(t, Len(Constant("abc"))).Eval(nil))
	assert.Equal(t, 1, msgpacked(t, Len(Constant(0))).Eval(nil))
	assert.Nil(t, msgpacked(t, Len(Constant(nil))).Eval(nil))
}
