package goexpr

import (
	"fmt"
)

// Substr takes a substring of the given source starting at the given index
// capped to the given length.
func Substr(source Expr, from Expr, length Expr) Expr {
	return &substr{source, from, length}
}

type substr struct {
	source Expr
	from   Expr
	length Expr
}

func (e *substr) Eval(params Params) interface{} {
	source := e.source.Eval(params)
	if source == nil {
		return nil
	}
	result := source.(string)
	from := e.from.Eval(params).(int)
	if from >= len(result) {
		return nil
	}
	result = result[from:]
	length := e.length.Eval(params).(int)
	if length > 0 && length < len(result) {
		result = result[:length]
	}
	return result
}

func (e *substr) WalkParams(cb func(string)) {
	e.source.WalkParams(cb)
	e.from.WalkParams(cb)
	e.length.WalkParams(cb)
}

func (e *substr) WalkOneToOneParams(cb func(string)) {
	// this function is not one-to-one, stop
}

func (e *substr) WalkLists(cb func(List)) {
	e.source.WalkLists(cb)
	e.from.WalkLists(cb)
	e.length.WalkLists(cb)
}

func (e *substr) String() string {
	return fmt.Sprintf("substr(%v,%v,%v)", e.source.String(), e.from.String(), e.length.String())
}
