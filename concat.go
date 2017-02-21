package goexpr

import (
	"bytes"
	"fmt"
)

// Concat joins a list of values using the first as a delimiter.
func Concat(exprs ...Expr) Expr {
	return &concat{exprs[0], exprs[1:]}
}

type concat struct {
	delim   Expr
	wrapped []Expr
}

func (e *concat) Eval(params Params) interface{} {
	delim := e.delim.Eval(params)
	buf := &bytes.Buffer{}
	for i, wrapped := range e.wrapped {
		first := i == 0
		if !first {
			fmt.Fprint(buf, delim)
		}
		fmt.Fprint(buf, wrapped.Eval(params))
	}
	return buf.String()
}

func (e *concat) WalkParams(cb func(string)) {
	e.delim.WalkParams(cb)
	for _, wrapped := range e.wrapped {
		wrapped.WalkParams(cb)
	}
}

func (e *concat) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *concat) WalkLists(cb func(List)) {
	e.delim.WalkLists(cb)
	for _, wrapped := range e.wrapped {
		wrapped.WalkLists(cb)
	}
}

func (e *concat) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("CONCAT(")
	buf.WriteString(e.delim.String())
	for _, wrapped := range e.wrapped {
		buf.WriteString(", ")
		fmt.Fprint(buf, wrapped.String())
	}
	buf.WriteString(")")
	return buf.String()
}
