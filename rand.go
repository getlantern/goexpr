package goexpr

import (
	"fmt"
	"math/rand"
)

// Rand randomly returns true, with ratio determining how frequently it returns
// true. A ratio of 0 will never return true, a ratio of 1 will always return
// true. If ratio fails to evaluate to anything, Rand returns false.
func Rand(ratio Expr) Expr {
	return &randExpr{ratio}
}

type randExpr struct {
	ratio Expr
}

func (e *randExpr) Eval(params Params) interface{} {
	ratio, ok := e.ratio.Eval(params).(float64)
	if !ok {
		return false
	}
	return rand.Float64() < ratio
}

func (e *randExpr) String() string {
	return fmt.Sprintf("RAND(%v)", e.ratio)
}
