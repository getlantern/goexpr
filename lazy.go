package goexpr

import (
	"fmt"
)

func Lazy(init func() interface{}) Expr {
	return &lazy{init: init}
}

type lazy struct {
	init        func() interface{}
	initialized bool
	val         interface{}
}

func (e *lazy) Eval(params Params) interface{} {
	if !e.initialized {
		e.val = e.init()
		e.initialized = true
	}
	return e.val
}

func (e *lazy) String() string {
	return fmt.Sprintf("LAZY(%v)", e.init)
}
