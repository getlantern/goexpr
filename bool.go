package goexpr

import (
	"fmt"
)

func Boolean(operator string, left Expr, right Expr) (Expr, error) {
	var bfn boolFN
	switch operator {
	case "AND":
		bfn = and
	case "OR":
		bfn = or
	default:
		return nil, fmt.Errorf("Unknown boolean operator %v", operator)
	}
	return &booleanExpr{operator, bfn, left, right}, nil
}

type boolFN func(a Expr, b Expr, params Params) bool

func and(a Expr, b Expr, params Params) bool {
	aVal, aok := a.Eval(params).(bool)
	if !aok || !aVal {
		return false
	}
	bVal, bok := b.Eval(params).(bool)
	return bok && bVal
}

func or(a Expr, b Expr, params Params) bool {
	aVal, aok := a.Eval(params).(bool)
	bVal, bok := b.Eval(params).(bool)
	return (aok && aVal) || (bok && bVal)
}

type booleanExpr struct {
	operatorString string
	operator       boolFN
	left           Expr
	right          Expr
}

func (e *booleanExpr) String() string {
	return fmt.Sprintf("(%v %v %v)", e.left, e.operatorString, e.right)
}

func (e *booleanExpr) Eval(params Params) interface{} {
	return e.operator(e.left, e.right, params)
}
