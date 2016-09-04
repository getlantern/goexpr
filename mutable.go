package goexpr

import (
	"fmt"
)

func Mutable() *MutableExpr {
	return &MutableExpr{}
}

type MutableExpr struct {
	wrapped Expr
}

func (e *MutableExpr) Set(wrapped Expr) {
	e.wrapped = wrapped
}

func (e *MutableExpr) Eval(params Params) interface{} {
	if e.wrapped == nil {
		return nil
	}
	return e.wrapped.Eval(params)
}

func (e *MutableExpr) String() string {
	if e.wrapped == nil {
		return "Mutable(<unset>)"
	}
	return fmt.Sprintf("Mutable(%v)", e.wrapped.String())
}
