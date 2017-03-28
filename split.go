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
	Source Expr
	Delim  Expr
	Idx    Expr
}

func (e *split) Eval(params Params) interface{} {
	source := e.Source.Eval(params)
	if source == nil {
		return nil
	}
	delim := e.Delim.Eval(params)
	idx := e.Idx.Eval(params).(int)
	parts := strings.Split(source.(string), delim.(string))
	if idx >= len(parts) {
		return nil
	}
	return parts[idx]
}

func (e *split) WalkParams(cb func(string)) {
	e.Source.WalkParams(cb)
	e.Delim.WalkParams(cb)
	e.Idx.WalkParams(cb)
}

func (e *split) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *split) WalkLists(cb func(List)) {
	e.Source.WalkLists(cb)
	e.Delim.WalkLists(cb)
	e.Idx.WalkLists(cb)
}

func (e *split) String() string {
	return fmt.Sprintf("SPLIT(%v,%v,%v)", e.Source.String(), e.Delim.String(), e.Idx.String())
}
