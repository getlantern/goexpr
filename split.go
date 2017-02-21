package goexpr

import (
	"fmt"
	"strings"
)

// Split splits a given source on a delimiter and returns the value at the given
// index. If no value exists at the given index, returns nil.
func Split(source Expr, delim Expr, idx Expr) Expr {
	return &split{source, delim, idx}
}

type split struct {
	source Expr
	delim  Expr
	idx    Expr
}

func (e *split) Eval(params Params) interface{} {
	source := e.source.Eval(params)
	if source == nil {
		return nil
	}
	delim := e.delim.Eval(params)
	idx := e.idx.Eval(params).(int)
	parts := strings.Split(source.(string), delim.(string))
	if idx >= len(parts) {
		return nil
	}
	return parts[idx]
}

func (e *split) WalkParams(cb func(string)) {
	e.source.WalkParams(cb)
	e.delim.WalkParams(cb)
	e.idx.WalkParams(cb)
}

func (e *split) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *split) WalkLists(cb func(List)) {
	e.source.WalkLists(cb)
	e.delim.WalkLists(cb)
	e.idx.WalkLists(cb)
}

func (e *split) String() string {
	return fmt.Sprintf("SPLIT(%v,%v,%v)", e.source.String(), e.delim.String(), e.idx.String())
}
