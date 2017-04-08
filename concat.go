package goexpr

import (
	"bytes"
	"fmt"
)

// Concat joins a list of values using the first as a delimiter.
func Concat(exprs ...Expr) Expr {
	return &concat{exprs[0], exprs[1:], false}
}

// PConcat joins a list of values using the first as a delimiter. Unlike Concat,
// PConcat assumes that it is a one-to-one function.
func PConcat(exprs ...Expr) Expr {
	return &concat{exprs[0], exprs[1:], true}
}

type concat struct {
	Delim    Expr
	Wrapped  []Expr
	OneToOne bool
}

func (e *concat) Eval(params Params) interface{} {
	delim := e.Delim.Eval(params)
	buf := &bytes.Buffer{}
	for i, wrapped := range e.Wrapped {
		first := i == 0
		if !first {
			fmt.Fprint(buf, delim)
		}
		val := wrapped.Eval(params)
		if val == nil {
			// replace nil with empty string
			val = ""
		}
		fmt.Fprint(buf, val)
	}
	return buf.String()
}

func (e *concat) WalkParams(cb func(string)) {
	e.Delim.WalkParams(cb)
	for _, wrapped := range e.Wrapped {
		wrapped.WalkParams(cb)
	}
}

func (e *concat) WalkOneToOneParams(cb func(string)) {
	if e.OneToOne {
		e.Delim.WalkOneToOneParams(cb)
		for _, wrapped := range e.Wrapped {
			wrapped.WalkOneToOneParams(cb)
		}
	}
}

func (e *concat) WalkLists(cb func(List)) {
	e.Delim.WalkLists(cb)
	for _, wrapped := range e.Wrapped {
		wrapped.WalkLists(cb)
	}
}

func (e *concat) String() string {
	buf := &bytes.Buffer{}
	if e.OneToOne {
		buf.WriteString("PCONCAT(")
	} else {
		buf.WriteString("CONCAT(")
	}
	buf.WriteString(e.Delim.String())
	for _, wrapped := range e.Wrapped {
		buf.WriteString(", ")
		fmt.Fprint(buf, wrapped.String())
	}
	buf.WriteString(")")
	return buf.String()
}
