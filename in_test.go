package goexpr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIn(t *testing.T) {
	list1 := ArrayList{Constant(1), Constant(3), Constant(6)}
	e1 := msgpacked(t, In(Constant(5), list1))
	e2 := msgpacked(t, In(Constant(5), ArrayList{Constant(1), Constant(3), Constant(5.0), Constant(6)}))
	assert.Equal(t, false, e1.Eval(nil))
	assert.Equal(t, true, e2.Eval(nil))
	assert.Equal(t, "5 IN(1, 3, 6)", e1.String())
	assert.Equal(t, "5 IN(1, 3, 5, 6)", e2.String())
	var lists []List
	e1.WalkLists(func(list List) {
		lists = append(lists, list)
	})
	assert.Len(t, lists, 1)
	assert.EqualValues(t, list1, lists[0])
}
