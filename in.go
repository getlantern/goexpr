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
	buf.WriteString(e.val.String())
	buf.WriteString(" IN(")
	for i, candidate := range e.candidates {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprint(buf, candidate.String())
	}
	buf.WriteString(")")
	return buf.String()
}
