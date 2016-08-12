package goexpr

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	zeroTime = time.Time{}
)

func Binary(operator string, left Expr, right Expr) (Expr, error) {
	// Normalize equals and not equals
	switch operator {
	case "=":
		operator = "=="
	case "<>":
		operator = "!="
	}
	return &binaryExpr{operator, buildOp(operator), left, right}, nil
}

type binaryExpr struct {
	operatorString string
	operator       func(a interface{}, b interface{}) interface{}
	left           Expr
	right          Expr
}

func (e *binaryExpr) String() string {
	return fmt.Sprintf("(%v %v %v)", e.left, e.operatorString, e.right)
}

func (e *binaryExpr) Eval(params Params) interface{} {
	return e.operator(e.left.Eval(params), e.right.Eval(params))
}

func buildOp(_oper string) func(interface{}, interface{}) interface{} {
	oper := ops[_oper]
	return func(a interface{}, b interface{}) interface{} {
		switch x := a.(type) {
		case nil:
			switch y := b.(type) {
			case nil:
				return oper.nl(x, y)
			case bool:
				return oper.bl(false, y)
			case byte:
				return oper.uin(0, uint64(y))
			case uint16:
				return oper.uin(0, uint64(y))
			case uint32:
				return oper.uin(0, uint64(y))
			case uint64:
				return oper.uin(0, uint64(y))
			case uint:
				return oper.uin(0, uint64(y))
			case int8:
				return oper.sin(0, int64(y))
			case int16:
				return oper.sin(0, int64(y))
			case int32:
				return oper.sin(0, int64(y))
			case int64:
				return oper.sin(0, int64(y))
			case int:
				return oper.sin(0, int64(y))
			case float32:
				return oper.fl(0, float64(y))
			case float64:
				return oper.fl(0, float64(y))
			case string:
				return oper.st("", y)
			case time.Time:
				return oper.ts(zeroTime, y)
			default:
				return oper.dflt
			}
		default:
			return oper.dflt
		}
	}
}

type op struct {
	nl   func(a interface{}, b interface{}) interface{}
	bl   func(a bool, b bool) interface{}
	uin  func(a uint64, b uint64) interface{}
	sin  func(a int64, b int64) interface{}
	fl   func(a float64, b float64) interface{}
	st   func(a string, b string) interface{}
	ts   func(a time.Time, b time.Time) interface{}
	dflt interface{}
}

var ops = map[string]op{
	"==": op{
		nl: func(a interface{}, b interface{}) interface{} {
			if a == nil {
				return b == nil
			}
			return true
		},
		bl: func(a bool, b bool) interface{} {
			return a == b
		},
		uin: func(a uint64, b uint64) interface{} {
			return a == b
		},
		sin: func(a int64, b int64) interface{} {
			return a == b
		},
		fl: func(a float64, b float64) interface{} {
			return a == b
		},
		st: func(a string, b string) interface{} {
			return a == b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return a == b
		},
		dflt: false,
	},
	"LIKE": op{
		nl: func(a interface{}, b interface{}) interface{} {
			if a == nil {
				return b == nil
			}
			return true
		},
		bl: func(a bool, b bool) interface{} {
			return a == b
		},
		uin: func(a uint64, b uint64) interface{} {
			return a == b
		},
		sin: func(a int64, b int64) interface{} {
			return a == b
		},
		fl: func(a float64, b float64) interface{} {
			return a == b
		},
		st: func(a string, b string) interface{} {
			lb := len(b)
			last := lb - 1
			startWildcard := b[0] == '%'
			endWildcard := b[last] == '%'
			if endWildcard {
				b = b[:last]
			}
			if startWildcard {
				b = b[1:]
			}
			if lb == 0 {
				return true
			}
			if startWildcard && endWildcard {
				return strings.Contains(a, b)
			}
			la := len(a)
			if la < lb {
				return false
			}
			if endWildcard {
				return a[:lb] == b
			}
			return a[lb-la:] == b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return a == b
		},
		dflt: false,
	},
	"!=": op{
		nl: func(a interface{}, b interface{}) interface{} {
			if a == nil {
				return b != nil
			}
			return b == nil
		},
		bl: func(a bool, b bool) interface{} {
			return a != b
		},
		uin: func(a uint64, b uint64) interface{} {
			return a != b
		},
		sin: func(a int64, b int64) interface{} {
			return a != b
		},
		fl: func(a float64, b float64) interface{} {
			return a != b
		},
		st: func(a string, b string) interface{} {
			return a != b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return a != b
		},
		dflt: false,
	},
	"<": op{
		nl: func(a interface{}, b interface{}) interface{} {
			return a == nil && b != nil
		},
		bl: func(a bool, b bool) interface{} {
			return !a && b
		},
		uin: func(a uint64, b uint64) interface{} {
			return a < b
		},
		sin: func(a int64, b int64) interface{} {
			return a < b
		},
		fl: func(a float64, b float64) interface{} {
			return a < b
		},
		st: func(a string, b string) interface{} {
			return a < b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return a.Before(b)
		},
		dflt: false,
	},
	"<=": op{
		nl: func(a interface{}, b interface{}) interface{} {
			return a == nil
		},
		bl: func(a bool, b bool) interface{} {
			return true
		},
		uin: func(a uint64, b uint64) interface{} {
			return a <= b
		},
		sin: func(a int64, b int64) interface{} {
			return a <= b
		},
		fl: func(a float64, b float64) interface{} {
			return a <= b
		},
		st: func(a string, b string) interface{} {
			return a <= b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return !a.After(b)
		},
		dflt: false,
	},
	">": op{
		nl: func(a interface{}, b interface{}) interface{} {
			return a != nil && b == nil
		},
		bl: func(a bool, b bool) interface{} {
			return a && !b
		},
		uin: func(a uint64, b uint64) interface{} {
			return a > b
		},
		sin: func(a int64, b int64) interface{} {
			return a > b
		},
		fl: func(a float64, b float64) interface{} {
			return a > b
		},
		st: func(a string, b string) interface{} {
			return a > b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return a.After(b)
		},
		dflt: false,
	},
	">=": op{
		nl: func(a interface{}, b interface{}) interface{} {
			return b == nil
		},
		bl: func(a bool, b bool) interface{} {
			return true
		},
		uin: func(a uint64, b uint64) interface{} {
			return a >= b
		},
		sin: func(a int64, b int64) interface{} {
			return a >= b
		},
		fl: func(a float64, b float64) interface{} {
			return a >= b
		},
		st: func(a string, b string) interface{} {
			return a >= b
		},
		ts: func(a time.Time, b time.Time) interface{} {
			return !a.Before(b)
		},
		dflt: false,
	},
	"+": op{
		uin: func(a uint64, b uint64) interface{} {
			return a + b
		},
		sin: func(a int64, b int64) interface{} {
			return a + b
		},
		fl: func(a float64, b float64) interface{} {
			return a + b
		},
		dflt: 0,
	},
	"-": op{
		uin: func(a uint64, b uint64) interface{} {
			return a - b
		},
		sin: func(a int64, b int64) interface{} {
			return a - b
		},
		fl: func(a float64, b float64) interface{} {
			return a - b
		},
		dflt: 0,
	},
	"*": op{
		uin: func(a uint64, b uint64) interface{} {
			return a * b
		},
		sin: func(a int64, b int64) interface{} {
			return a * b
		},
		fl: func(a float64, b float64) interface{} {
			return a * b
		},
		dflt: 0,
	},
	"/": op{
		uin: func(a uint64, b uint64) interface{} {
			if b == 0 {
				return uint64(0)
			}
			return a / b
		},
		sin: func(a int64, b int64) interface{} {
			if b == 0 {
				return int64(0)
			}
			return a / b
		},
		fl: func(a float64, b float64) interface{} {
			if b == 0 {
				return 0
			}
			return a / b
		},
		dflt: 0,
	},
}

func strToBool(str string) (bool, bool) {
	r, err := strconv.ParseBool(str)
	return r, err == nil
}

func strToUint(str string) (uint64, bool) {
	r, err := strconv.ParseUint(str, 10, 64)
	return r, err == nil
}

func strToInt(str string) (int64, bool) {
	r, err := strconv.ParseInt(str, 10, 64)
	return r, err == nil
}

func strToFloat(str string) (float64, bool) {
	r, err := strconv.ParseFloat(str, 64)
	return r, err == nil
}

func timeToBool(t time.Time) bool {
	return !t.IsZero()
}

func div(x interface{}, y interface{}) interface{} {
	switch v1 := x.(type) {
	case uint64:
		v2 := y.(uint64)
		if y == 0 {
			return 0
		}
		return v1 + v2
	case int64:
		v2 := y.(int64)
		if y == 0 {
			return 0
		}
		return v1 + v2
	case float64:
		v2 := y.(float64)
		if y == 0 {
			return 0
		}
		return v1 + v2
	}
	return 0
}
