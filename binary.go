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
	case "==", "<", ">", "LIKE", "+", "-", "*", "/", "OR", "AND":
		o = ops[operator]
	case "!=":
		o = not(ops["=="])
	case "<=":
		o = or(ops["<"], ops["=="])
	case ">=":
		o = or(ops[">"], ops["=="])
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
		switch v2 := b.(type) {
		case bool:
			return v == v2
		default:
			return false
		}
	case byte:
		switch v2 := b.(type) {
		case byte:
			return v == v2
		default:
			return false
		}
	case uint16:
		switch v2 := b.(type) {
		case uint16:
			return v == v2
		default:
			return false
		}
	case uint32:
		switch v2 := b.(type) {
		case uint32:
			return v == v2
		default:
			return false
		}
	case uint64:
		switch v2 := b.(type) {
		case uint64:
			return v == v2
		default:
			return false
		}
	case int8:
		switch v2 := b.(type) {
		case int8:
			return v == v2
		default:
			return false
		}
	case int16:
		switch v2 := b.(type) {
		case int16:
			return v == v2
		default:
			return false
		}
	case int32:
		switch v2 := b.(type) {
		case int32:
			return v == v2
		default:
			return false
		}
	case int64:
		switch v2 := b.(type) {
		case int64:
			return v == v2
		default:
			return false
		}
	case int:
		switch v2 := b.(type) {
		case int:
			return v == v2
		default:
			return false
		}
	case float32:
		switch v2 := b.(type) {
		case float32:
			return v == v2
		default:
			return false
		}
	case float64:
		switch v2 := b.(type) {
		case float64:
			return v == v2
		default:
			return false
		}
	case string:
		switch v2 := b.(type) {
		case string:
			return v == v2
		default:
			return false
		}
	case time.Time:
		switch v2 := b.(type) {
		case time.Time:
			return v == v2
		default:
			return false
		}
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
			switch v2 := b.(type) {
			case bool:
				return !v && v2
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case byte:
				return v < v2
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case uint16:
				return v < v2
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				return v < v2
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				return v < v2
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				return v < v2
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				return v < v2
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				return v < v2
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				return v < v2
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case int:
				return v < v2
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				return v < v2
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				return v < v2
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case string:
				return v < v2
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case time.Time:
				return v.Before(v2)
			default:
				return false
			}
		default:
			return false
		}
	},
	">": func(a interface{}, b interface{}) interface{} {
		if a == nil {
			return false
		}
		if b == nil {
			return true
		}
		switch v := a.(type) {
		case bool:
			switch v2 := b.(type) {
			case bool:
				return v && !v2
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case byte:
				return v > v2
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case uint16:
				return v > v2
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				return v > v2
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				return v > v2
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				return v > v2
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				return v > v2
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				return v > v2
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				return v > v2
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case int:
				return v > v2
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				return v > v2
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				return v > v2
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case string:
				return v > v2
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case time.Time:
				return v.After(v2)
			default:
				return false
			}
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
			switch v2 := b.(type) {
			case string:
				return strings.Contains(v, v2)
			default:
				return strings.Contains(v, fmt.Sprint(b))
			}
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
			switch v2 := b.(type) {
			case uint16:
				return v + v2
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				return v + v2
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				return v + v2
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				return v + v2
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				return v + v2
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				return v + v2
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				return v + v2
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case int:
				return v + v2
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				return v + v2
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				return v + v2
			default:
				return 0
			}
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
			switch v2 := b.(type) {
			case uint16:
				return v - v2
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				return v - v2
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				return v - v2
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				return v - v2
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				return v - v2
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				return v - v2
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				return v - v2
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case int:
				return v - v2
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				return v - v2
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				return v - v2
			default:
				return 0
			}
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
			switch v2 := b.(type) {
			case uint16:
				return v * v2
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				return v * v2
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				return v * v2
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				return v * v2
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				return v * v2
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				return v * v2
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				return v * v2
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case int:
				return v * v2
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				return v * v2
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				return v * v2
			default:
				return 0
			}
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
			switch v2 := b.(type) {
			case uint16:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case uint32:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case uint64:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case int8:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case int16:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case int32:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case int64:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case int:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case float32:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case float64:
				if v2 == 0 {
					return 0
				}
				return v / v2
			default:
				return 0
			}
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
