package goexpr

import (
	"bytes"
	"fmt"
)

func In(val Expr, candidates ...Expr) Expr {
	return &in{val, candidates}
}

type in struct {
	val        Expr
	candidates []Expr
}

func (e *in) Eval(params Params) interface{} {
	v := e.val.Eval(params)
	for _, candidate := range e.candidates {
		c := candidate.Eval(params)
		if c == v {
			return true
		}
	}
	return false
}

func (e *in) String() string {
	buf := &bytes.Buffer{}
	buf.WriteString("IN(")
	buf.WriteString(e.val.String())
	for _, candidate := range e.candidates {
		buf.WriteString(", ")
		fmt.Fprint(buf, candidate.String())
	}
	buf.WriteString(")")
	return buf.String()
}
