package goexpr

import (
	"fmt"
	"strings"
	"time"
)

type op func(a interface{}, b interface{}) interface{}

func Binary(operator string, left Expr, right Expr) (Expr, error) {
	// Normalize equals and not equals
	switch operator {
	case "=":
		operator = "=="
	case "<>":
		operator = "!="
	}

	// Some operators are just the logical negation or combination of different
	// operators, use that.
	var o op
	switch operator {
	case "==", "<", "LIKE", "+", "-", "*", "/", "OR", "AND":
		o = ops[operator]
	case "!=":
		o = not(ops["=="])
	case "<=":
		o = or(ops["<"], ops["=="])
	case ">=":
		o = not(ops["<"])
	case ">":
		o = not(or(ops["<"], ops["=="]))
	case "NOT LIKE":
		o = not(ops["LIKE"])
	default:
		return nil, fmt.Errorf("Unknown operator %v", operator)
	}
	return &binaryExpr{operator, o, left, right}, nil
}

type binaryExpr struct {
	operatorString string
	operator       op
	left           Expr
	right          Expr
}

func (e *binaryExpr) String() string {
	return fmt.Sprintf("(%v %v %v)", e.left, e.operatorString, e.right)
}

func not(operator op) op {
	return func(a interface{}, b interface{}) interface{} {
		return !operator(a, b).(bool)
	}
}

func or(o1 op, o2 op) op {
	return func(a interface{}, b interface{}) interface{} {
		if o1(a, b).(bool) {
			return true
		}
		return o2(a, b).(bool)
	}
}

func (e *binaryExpr) Eval(params Params) interface{} {
	return e.operator(e.left.Eval(params), e.right.Eval(params))
}

var eq = func(a interface{}, b interface{}) interface{} {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	switch v := a.(type) {
	case bool:
		return v == b.(bool)
	case byte:
		return v == b.(byte)
	case uint16:
		return v == b.(uint16)
	case uint32:
		return v == b.(uint32)
	case uint64:
		return v == b.(uint64)
	case int8:
		return v == b.(int8)
	case int16:
		return v == b.(int16)
	case int32:
		return v == b.(int32)
	case int64:
		return v == b.(int64)
	case int:
		return v == b.(int)
	case float32:
		return v == b.(float32)
	case float64:
		return v == b.(float64)
	case string:
		return v == b.(string)
	case time.Time:
		return v == b.(time.Time)
	default:
		return false
	}
}

var ops = map[string]op{
	"==": eq,
	"<": func(a interface{}, b interface{}) interface{} {
		if b == nil {
			return false
		}
		if a == nil {
			return true
		}
		switch v := a.(type) {
		case bool:
			return !v && b.(bool)
		case byte:
			return v < b.(byte)
		case uint16:
			return v < b.(uint16)
		case uint32:
			return v < b.(uint32)
		case uint64:
			return v < b.(uint64)
		case int8:
			return v < b.(int8)
		case int16:
			return v < b.(int16)
		case int32:
			return v < b.(int32)
		case int64:
			return v < b.(int64)
		case int:
			return v < b.(int)
		case float32:
			return v < b.(float32)
		case float64:
			return v < b.(float64)
		case string:
			return v < b.(string)
		case time.Time:
			return v.Before(b.(time.Time))
		default:
			return false
		}
	},
	"LIKE": func(a interface{}, b interface{}) interface{} {
		if a == nil {
			return b == nil
		}
		if b == nil {
			return false
		}
		switch v := a.(type) {
		case string:
			return strings.Contains(v, b.(string))
		default:
			return eq(a, b)
		}
	},
	"+": func(a interface{}, b interface{}) interface{} {
		if a == nil || b == nil {
			return 0
		}
		switch v := a.(type) {
		case uint16:
			return v + b.(uint16)
		case uint32:
			return v + b.(uint32)
		case uint64:
			return v + b.(uint64)
		case int8:
			return v + b.(int8)
		case int16:
			return v + b.(int16)
		case int32:
			return v + b.(int32)
		case int64:
			return v + b.(int64)
		case int:
			return v + b.(int)
		case float32:
			return v + b.(float32)
		case float64:
			return v + b.(float64)
		default:
			return 0
		}
	},
	"-": func(a interface{}, b interface{}) interface{} {
		if a == nil || b == nil {
			return 0
		}
		switch v := a.(type) {
		case uint16:
			return v - b.(uint16)
		case uint32:
			return v - b.(uint32)
		case uint64:
			return v - b.(uint64)
		case int8:
			return v - b.(int8)
		case int16:
			return v - b.(int16)
		case int32:
			return v - b.(int32)
		case int64:
			return v - b.(int64)
		case int:
			return v - b.(int)
		case float32:
			return v - b.(float32)
		case float64:
			return v - b.(float64)
		default:
			return 0
		}
	},
	"*": func(a interface{}, b interface{}) interface{} {
		if a == nil || b == nil {
			return 0
		}
		switch v := a.(type) {
		case uint16:
			return v * b.(uint16)
		case uint32:
			return v * b.(uint32)
		case uint64:
			return v * b.(uint64)
		case int8:
			return v * b.(int8)
		case int16:
			return v * b.(int16)
		case int32:
			return v * b.(int32)
		case int64:
			return v * b.(int64)
		case int:
			return v * b.(int)
		case float32:
			return v * b.(float32)
		case float64:
			return v * b.(float64)
		default:
			return 0
		}
	},
	"/": func(a interface{}, b interface{}) interface{} {
		if a == nil || b == nil {
			return 0
		}
		switch v := a.(type) {
		case uint16:
			w := b.(uint16)
			if b == 0 {
				return 0
			}
			return v / w
		case uint32:
			w := b.(uint32)
			if b == 0 {
				return 0
			}
			return v / w
		case uint64:
			w := b.(uint64)
			if b == 0 {
				return 0
			}
			return v / w
		case int8:
			w := b.(int8)
			if b == 0 {
				return 0
			}
			return v / w
		case int16:
			w := b.(int16)
			if b == 0 {
				return 0
			}
			return v / w
		case int32:
			w := b.(int32)
			if b == 0 {
				return 0
			}
			return v / w
		case int64:
			w := b.(int64)
			if b == 0 {
				return 0
			}
			return v / w
		case int:
			w := b.(int)
			if b == 0 {
				return 0
			}
			return v / w
		case float32:
			w := b.(float32)
			if b == 0 {
				return 0
			}
			return v / w
		case float64:
			w := b.(float64)
			if b == 0 {
				return 0
			}
			return v / w
		default:
			return 0
		}
	},
	"AND": func(a interface{}, b interface{}) interface{} {
		return a.(bool) && b.(bool)
	},
	"OR": func(a interface{}, b interface{}) interface{} {
		return a.(bool) || b.(bool)
	},
}
