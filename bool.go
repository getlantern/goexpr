package goexpr

import (
	"fmt"
)

// Boolean accepts the operators AND, OR and returns a short-circuiting
// expression that evaluates left first and right second.
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
	if aok && aVal {
		return true
	}
	bVal, bok := b.Eval(params).(bool)
	return (aok && aVal) || (bok && bVal)
}

type booleanExpr struct {
	operatorString string
	operator       boolFN
	left           Expr
	right          Expr
}

func (e *booleanExpr) Eval(params Params) interface{} {
	return e.operator(e.left, e.right, params)
}

func (e *booleanExpr) WalkParams(cb func(string)) {
	e.left.WalkParams(cb)
	e.right.WalkParams(cb)
}

func (e *booleanExpr) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *booleanExpr) WalkLists(cb func(List)) {
	e.left.WalkLists(cb)
	e.right.WalkLists(cb)
}

func (e *booleanExpr) String() string {
	return fmt.Sprintf("(%v %v %v)", e.left, e.operatorString, e.right)
}
