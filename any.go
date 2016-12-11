package goexpr

import (
	"fmt"
)

// Any is an expression that returns the first non-nil, non-empty string value
// from evaluating the expressions candidates.
func Any(candidates List) Expr {
	return &any{_candidates: candidates}
}

type any struct {
	_candidates List
	initialized bool
	candidates  []Expr
}

func (e *any) Eval(params Params) interface{} {
	if !e.initialized {
		e.candidates = e._candidates.Values()
		e.initialized = true
	}
	for _, candidate := range e.candidates {
		c := candidate.Eval(params)
		if c != nil && c != "" {
			return c
		}
	}
	return nil
}

func (e *any) WalkLists(cb func(List)) {
	cb(e._candidates)
}

func (e *any) String() string {
	return fmt.Sprintf("ANY(%v)", e._candidates)
}
