package goexpr

import (
	"strings"
	"time"
)

var _ops = map[string]func(interface{}, interface{}) interface{}{
	"==": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return true
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case bool:
				x, y := v1, v2
				return x == y
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return x == y
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return x == y
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				x, y := v1, v2
				return x == y
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return x == y
			case time.Time:
				x, y := v1, v2
				return x == y
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				x, y := v1, v2
				return x == y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				x, y := v1, v2
				return x == y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				x, y := v1, v2
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				x, y := v1, v2
				return x == y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				x, y := v1, v2
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				x, y := v1, v2
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				x, y := v1, v2
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				x, y := v1, v2
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				x, y := v1, v2
				return x == y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				x, y := v1, v2
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				x, y := v1, v2
				return x == y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x == y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x == y
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x == y
			case float64:
				x, y := v1, v2
				return x == y
			default:
				return false
			}
		default:
			return false
		}
	},
	"!=": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return false
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return x == y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case bool:
				x, y := v1, v2
				return x != y
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return x != y
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return x != y
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return x == y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				x, y := v1, v2
				return x == y
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return x == y
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return x == y
			case time.Time:
				x, y := v1, v2
				return x != y
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				x, y := v1, v2
				return x != y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				x, y := v1, v2
				return x != y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				x, y := v1, v2
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				x, y := v1, v2
				return x != y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				x, y := v1, v2
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				x, y := v1, v2
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				x, y := v1, v2
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				x, y := v1, v2
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				x, y := v1, v2
				return x != y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				x, y := v1, v2
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				x, y := v1, v2
				return x != y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x != y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x != y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x != y
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x != y
			case float64:
				x, y := v1, v2
				return x != y
			default:
				return false
			}
		default:
			return false
		}
	},
	"<": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return false
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) < 0
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case bool:
				x, y := v1, v2
				return !x && y
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return !x && y
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return !x && y
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) < 0
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				x, y := v1, v2
				return strings.Compare(x, y) < 0
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) < 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) < 0
			case time.Time:
				x, y := v1, v2
				return x.Before(y)
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				x, y := v1, v2
				return x < y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				x, y := v1, v2
				return x < y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				x, y := v1, v2
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				x, y := v1, v2
				return x < y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				x, y := v1, v2
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				x, y := v1, v2
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				x, y := v1, v2
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				x, y := v1, v2
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				x, y := v1, v2
				return x < y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				x, y := v1, v2
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				x, y := v1, v2
				return x < y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x < y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return !x && y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.Before(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x < y
			case float64:
				x, y := v1, v2
				return x < y
			default:
				return false
			}
		default:
			return false
		}
	},
	"<=": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return true
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) <= 0
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return true
			case bool:
				x, y := v1, v2
				return true
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return true
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return true
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) <= 0
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return true
			case string:
				x, y := v1, v2
				return strings.Compare(x, y) <= 0
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) <= 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) <= 0
			case time.Time:
				x, y := v1, v2
				return !x.After(y)
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				x, y := v1, v2
				return x <= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				x, y := v1, v2
				return x <= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				x, y := v1, v2
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				x, y := v1, v2
				return x <= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				x, y := v1, v2
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				x, y := v1, v2
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				x, y := v1, v2
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				x, y := v1, v2
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				x, y := v1, v2
				return x <= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				x, y := v1, v2
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				x, y := v1, v2
				return x <= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.After(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x <= y
			case float64:
				x, y := v1, v2
				return x <= y
			default:
				return false
			}
		default:
			return false
		}
	},
	">": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return false
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) > 0
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case bool:
				x, y := v1, v2
				return x && !y
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return x && !y
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return x && !y
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) > 0
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				x, y := v1, v2
				return strings.Compare(x, y) > 0
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) > 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) > 0
			case time.Time:
				x, y := v1, v2
				return x.After(y)
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				x, y := v1, v2
				return x > y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				x, y := v1, v2
				return x > y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				x, y := v1, v2
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				x, y := v1, v2
				return x > y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				x, y := v1, v2
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				x, y := v1, v2
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				x, y := v1, v2
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				x, y := v1, v2
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				x, y := v1, v2
				return x > y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				x, y := v1, v2
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				x, y := v1, v2
				return x > y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x > y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return x && !y
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return x.After(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x > y
			case float64:
				x, y := v1, v2
				return x > y
			default:
				return false
			}
		default:
			return false
		}
	},
	">=": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				x, y := v1, v2
				return true
			case bool:
				var x bool
				var xok bool
				x, xok = false, true
				var y bool
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) >= 0
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Time{}, true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				var y bool
				var xok bool
				y, xok = false, true
				var x bool
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return false
				}
				return true
			case bool:
				x, y := v1, v2
				return true
			case string:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = strToBool(v2)
				if !xok || !yok {
					return false
				}
				return true
			case time.Time:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = timeToBool(v2), true
				if !xok || !yok {
					return false
				}
				return true
			case byte:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case uint:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int8:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int16:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case int:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case float32:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case float64:
				var x bool
				var xok bool
				x, xok = v1, true
				var y bool
				var yok bool
				y, yok = v2 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			default:
				return false
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) >= 0
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = strToBool(v1)
				if !xok || !yok {
					return false
				}
				return true
			case string:
				x, y := v1, v2
				return strings.Compare(x, y) >= 0
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) >= 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				var y time.Time
				var xok bool
				y, xok = time.Time{}, true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = timeToBool(v1), true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return strings.Compare(x, y) >= 0
			case time.Time:
				x, y := v1, v2
				return !x.Before(y)
			case byte:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case uint16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case uint32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case uint64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case uint:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case int8:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case int16:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case int32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case int64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case int:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case float32:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case float64:
				var y time.Time
				var xok bool
				y, xok = time.Unix(int64(v2), 0), true
				var x time.Time
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			default:
				return false
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				x, y := v1, v2
				return x >= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				x, y := v1, v2
				return x >= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				x, y := v1, v2
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				x, y := v1, v2
				return x >= y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				x, y := v1, v2
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				x, y := v1, v2
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				x, y := v1, v2
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				x, y := v1, v2
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				x, y := v1, v2
				return x >= y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				x, y := v1, v2
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				x, y := v1, v2
				return x >= y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			default:
				return false
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case bool:
				var y bool
				var xok bool
				y, xok = v2, true
				var x bool
				var yok bool
				x, yok = v1 == 1, true
				if !xok || !yok {
					return false
				}
				return true
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case time.Time:
				var x time.Time
				var xok bool
				x, xok = time.Unix(int64(v1), 0), true
				var y time.Time
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return zeroTime
				}
				return !x.Before(y)
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x >= y
			case float64:
				x, y := v1, v2
				return x >= y
			default:
				return false
			}
		default:
			return false
		}
	},
	"+": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				var x string
				var xok bool
				x, xok = "", true
				var y string
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return ""
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				var y string
				var xok bool
				y, xok = "", true
				var x string
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return ""
				}
				return x + y
			case bool:
				return 0
			case string:
				x, y := v1, v2
				return x + y
			case time.Time:
				var x string
				var xok bool
				x, xok = v1, true
				var y string
				var yok bool
				y, yok = v2.String(), true
				if !xok || !yok {
					return ""
				}
				return x + y
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				var y string
				var xok bool
				y, xok = v2, true
				var x string
				var yok bool
				x, yok = v1.String(), true
				if !xok || !yok {
					return ""
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				x, y := v1, v2
				return x + y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				x, y := v1, v2
				return x + y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				x, y := v1, v2
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				x, y := v1, v2
				return x + y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				x, y := v1, v2
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				x, y := v1, v2
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				x, y := v1, v2
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				x, y := v1, v2
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				x, y := v1, v2
				return x + y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				x, y := v1, v2
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				x, y := v1, v2
				return x + y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x + y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x + y
			case float64:
				x, y := v1, v2
				return x + y
			default:
				return 0
			}
		default:
			return 0
		}
	},
	"-": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				x, y := v1, v2
				return x - y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				x, y := v1, v2
				return x - y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				x, y := v1, v2
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				x, y := v1, v2
				return x - y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				x, y := v1, v2
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				x, y := v1, v2
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				x, y := v1, v2
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				x, y := v1, v2
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				x, y := v1, v2
				return x - y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				x, y := v1, v2
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				x, y := v1, v2
				return x - y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x - y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x - y
			case float64:
				x, y := v1, v2
				return x - y
			default:
				return 0
			}
		default:
			return 0
		}
	},
	"*": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				x, y := v1, v2
				return x * y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				x, y := v1, v2
				return x * y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				x, y := v1, v2
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				x, y := v1, v2
				return x * y
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				x, y := v1, v2
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				x, y := v1, v2
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				x, y := v1, v2
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				x, y := v1, v2
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				x, y := v1, v2
				return x * y
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				x, y := v1, v2
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				x, y := v1, v2
				return x * y
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return x * y
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return x * y
			case float64:
				x, y := v1, v2
				return x * y
			default:
				return 0
			}
		default:
			return 0
		}
	},
	"/": func(a interface{}, b interface{}) interface{} {
		switch v1 := a.(type) {
		case nil:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x byte
				var xok bool
				x, xok = 0, true
				var y byte
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var x uint16
				var xok bool
				x, xok = 0, true
				var y uint16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var x uint32
				var xok bool
				x, xok = 0, true
				var y uint32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = 0, true
				var y uint64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var x uint
				var xok bool
				x, xok = 0, true
				var y uint
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int8
				var xok bool
				x, xok = 0, true
				var y int8
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int16
				var xok bool
				x, xok = 0, true
				var y int16
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int32
				var xok bool
				x, xok = 0, true
				var y int32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = 0, true
				var y int64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int
				var xok bool
				x, xok = 0, true
				var y int
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float32
				var xok bool
				x, xok = 0, true
				var y float32
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = 0, true
				var y float64
				var yok bool
				y, yok = v2, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case bool:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case string:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var x uint64
				var xok bool
				x, xok = strToUint(v1)
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = strToInt(v1)
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = strToFloat(v1)
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case time.Time:
			switch v2 := b.(type) {
			case nil:
				return 0
			case bool:
				return 0
			case string:
				return 0
			case time.Time:
				return 0
			case byte:
				return 0
			case uint16:
				return 0
			case uint32:
				return 0
			case uint64:
				return 0
			case uint:
				return 0
			case int8:
				return 0
			case int16:
				return 0
			case int32:
				return 0
			case int64:
				return 0
			case int:
				return 0
			case float32:
				return 0
			case float64:
				return 0
			default:
				return 0
			}
		case byte:
			switch v2 := b.(type) {
			case nil:
				var y byte
				var xok bool
				y, xok = 0, true
				var x byte
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				x, y := v1, v2
				return div(x, y)
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case uint16:
			switch v2 := b.(type) {
			case nil:
				var y uint16
				var xok bool
				y, xok = 0, true
				var x uint16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				x, y := v1, v2
				return div(x, y)
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case uint32:
			switch v2 := b.(type) {
			case nil:
				var y uint32
				var xok bool
				y, xok = 0, true
				var x uint32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				x, y := v1, v2
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case uint64:
			switch v2 := b.(type) {
			case nil:
				var y uint64
				var xok bool
				y, xok = 0, true
				var x uint64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				x, y := v1, v2
				return div(x, y)
			case uint:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case uint:
			switch v2 := b.(type) {
			case nil:
				var y uint
				var xok bool
				y, xok = 0, true
				var x uint
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y uint64
				var xok bool
				y, xok = strToUint(v2)
				var x uint64
				var yok bool
				x, yok = uint64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var x uint64
				var xok bool
				x, xok = uint64(v1), true
				var y uint64
				var yok bool
				y, yok = uint64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				x, y := v1, v2
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case int8:
			switch v2 := b.(type) {
			case nil:
				var y int8
				var xok bool
				y, xok = 0, true
				var x int8
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				x, y := v1, v2
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case int16:
			switch v2 := b.(type) {
			case nil:
				var y int16
				var xok bool
				y, xok = 0, true
				var x int16
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				x, y := v1, v2
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case int32:
			switch v2 := b.(type) {
			case nil:
				var y int32
				var xok bool
				y, xok = 0, true
				var x int32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				x, y := v1, v2
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case int64:
			switch v2 := b.(type) {
			case nil:
				var y int64
				var xok bool
				y, xok = 0, true
				var x int64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				x, y := v1, v2
				return div(x, y)
			case int:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case int:
			switch v2 := b.(type) {
			case nil:
				var y int
				var xok bool
				y, xok = 0, true
				var x int
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y int64
				var xok bool
				y, xok = strToInt(v2)
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y int64
				var xok bool
				y, xok = int64(v2), true
				var x int64
				var yok bool
				x, yok = int64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var x int64
				var xok bool
				x, xok = int64(v1), true
				var y int64
				var yok bool
				y, yok = int64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				x, y := v1, v2
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case float32:
			switch v2 := b.(type) {
			case nil:
				var y float32
				var xok bool
				y, xok = 0, true
				var x float32
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				x, y := v1, v2
				return div(x, y)
			case float64:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			default:
				return 0
			}
		case float64:
			switch v2 := b.(type) {
			case nil:
				var y float64
				var xok bool
				y, xok = 0, true
				var x float64
				var yok bool
				x, yok = v1, true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case bool:
				return 0
			case string:
				var y float64
				var xok bool
				y, xok = strToFloat(v2)
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case time.Time:
				return 0
			case byte:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case uint:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int8:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int16:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int32:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int64:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case int:
				var y float64
				var xok bool
				y, xok = float64(v2), true
				var x float64
				var yok bool
				x, yok = float64(v1), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float32:
				var x float64
				var xok bool
				x, xok = float64(v1), true
				var y float64
				var yok bool
				y, yok = float64(v2), true
				if !xok || !yok {
					return 0
				}
				return div(x, y)
			case float64:
				x, y := v1, v2
				return div(x, y)
			default:
				return 0
			}
		default:
			return 0
		}
	},
}
