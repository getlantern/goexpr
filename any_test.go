package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAny(t *testing.T) {
	l1 := ArrayList{Constant(nil), Constant(""), Constant(1)}
	e1 := Any(l1)
	e2 := Any(ArrayList{Constant(2), Constant(nil), Constant("")})
	assert.Equal(t, 1, e1.Eval(nil))
	assert.Equal(t, 2, e2.Eval(nil))
	var lists []List
	e1.WalkLists(func(list List) {
		lists = append(lists, list)
	})
	assert.Len(t, lists, 1)
	assert.EqualValues(t, l1, lists[0])
}
