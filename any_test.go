package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	e1 := msgpacked(t, Any(Constant(nil), Constant(""), Constant(1)))
	e2 := msgpacked(t, Any(Constant(2), Constant(nil), Constant("")))
	assert.EqualValues(t, 1, e1.Eval(nil))
	assert.EqualValues(t, 2, e2.Eval(nil))
}
