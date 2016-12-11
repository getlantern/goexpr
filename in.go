package goexpr

import (
	"bytes"
	"fmt"
)

type List interface {
	Values() []Expr
}

type ArrayList []Expr

func (al ArrayList) Values() []Expr {
	return al
}

func (al ArrayList) String() string {
	buf := &bytes.Buffer{}
	for i, e := range al {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprint(buf, e.String())
	}
	return buf.String()
}

func In(val Expr, candidates List) Expr {
	return &in{val: val, _candidates: candidates}
}

type in struct {
	val         Expr
	_candidates List
	initialized bool
	candidates  []Expr
}

func (e *in) Eval(params Params) interface{} {
	v := e.val.Eval(params)
	if !e.initialized {
		e.candidates = e._candidates.Values()
		e.initialized = true
	}
	for _, candidate := range e.candidates {
		c := candidate.Eval(params)
		if c == v {
			return true
		}
	}
	return false
}

func (e *in) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *in) WalkLists(cb func(List)) {
	cb(e._candidates)
	e.val.WalkLists(cb)
}

func (e *in) String() string {
	return fmt.Sprintf("%v IN(%v)", e.val.String(), e._candidates)
}
