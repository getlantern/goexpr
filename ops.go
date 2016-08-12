package goexpr

import (
  "time"
)

func buildOp(_o string) func(interface{}, interface{}) interface{} {
	o := ops[_o]
	return func(a interface{}, b interface{}) interface{} {
switch x := a.(type) {
  case nil:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    return o.bl(false, y)
    case string:
    return o.st("", y)
    case time.Time:
    return o.ts(zeroTime, y)
    case byte:
    return o.uin(0, uint64(y))
    case uint16:
    return o.uin(0, uint64(y))
    case uint32:
    return o.uin(0, uint64(y))
    case uint64:
    return o.uin(0, uint64(y))
    case uint:
    return o.uin(0, uint64(y))
    case int8:
    return o.sin(0, int64(y))
    case int16:
    return o.sin(0, int64(y))
    case int32:
    return o.sin(0, int64(y))
    case int64:
    return o.sin(0, int64(y))
    case int:
    return o.sin(0, int64(y))
    case float32:
    return o.fl(0, float64(y))
    case float64:
    return o.fl(0, float64(y))
    default:
      return o.dflt
  }
  case bool:
  switch y := b.(type) {
    case nil:
    return o.bl(x, false)
    case bool:
    return o.bl(x, y)
    case string:
    bv, ok := strToBool(y)
if !ok {
	return o.dflt
}
return o.bl(x, bv)
    case time.Time:
    return o.bl(x, timeToBool(y))
    case byte:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case uint16:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case uint32:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case uint64:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case uint:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case int8:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case int16:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case int32:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case int64:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case int:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case float32:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    case float64:
    if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt
    default:
      return o.dflt
  }
  case string:
  switch y := b.(type) {
    case nil:
    return o.st(x, "")
    case bool:
    bv, ok := strToBool(x)
if !ok {
	return o.dflt
}
return o.bl(bv, y)
    case string:
    return o.st(x, y)
    case time.Time:
    return o.st(x, y.String())
    case byte:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case uint16:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case uint32:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case uint64:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case uint:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case int8:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case int16:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case int32:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case int64:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case int:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case float32:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    case float64:
    nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))
    default:
      return o.dflt
  }
  case time.Time:
  switch y := b.(type) {
    case nil:
    return o.ts(x, zeroTime)
    case bool:
    return o.bl(timeToBool(x), y)
    case string:
    return o.st(x.String(), y)
    case time.Time:
    return o.ts(x, y)
    case byte:
    return o.uin(timeToUint(x), uint64(y))
    case uint16:
    return o.uin(timeToUint(x), uint64(y))
    case uint32:
    return o.uin(timeToUint(x), uint64(y))
    case uint64:
    return o.uin(timeToUint(x), uint64(y))
    case uint:
    return o.uin(timeToUint(x), uint64(y))
    case int8:
    return o.sin(timeToInt(x), int64(y))
    case int16:
    return o.sin(timeToInt(x), int64(y))
    case int32:
    return o.sin(timeToInt(x), int64(y))
    case int64:
    return o.sin(timeToInt(x), int64(y))
    case int:
    return o.sin(timeToInt(x), int64(y))
    case float32:
    return o.fl(timeToFloat(x), float64(y))
    case float64:
    return o.fl(timeToFloat(x), float64(y))
    default:
      return o.dflt
  }
  case byte:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.uin(uint64(x), timeToUint(y))
    case byte:
    return o.uin(uint64(x), uint64(y))
    case uint16:
    return o.uin(uint64(x), uint64(y))
    case uint32:
    return o.uin(uint64(x), uint64(y))
    case uint64:
    return o.uin(uint64(x), uint64(y))
    case uint:
    return o.uin(uint64(x), uint64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case uint16:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.uin(uint64(x), timeToUint(y))
    case byte:
    return o.uin(uint64(x), uint64(y))
    case uint16:
    return o.uin(uint64(x), uint64(y))
    case uint32:
    return o.uin(uint64(x), uint64(y))
    case uint64:
    return o.uin(uint64(x), uint64(y))
    case uint:
    return o.uin(uint64(x), uint64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case uint32:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.uin(uint64(x), timeToUint(y))
    case byte:
    return o.uin(uint64(x), uint64(y))
    case uint16:
    return o.uin(uint64(x), uint64(y))
    case uint32:
    return o.uin(uint64(x), uint64(y))
    case uint64:
    return o.uin(uint64(x), uint64(y))
    case uint:
    return o.uin(uint64(x), uint64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case uint64:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.uin(uint64(x), timeToUint(y))
    case byte:
    return o.uin(uint64(x), uint64(y))
    case uint16:
    return o.uin(uint64(x), uint64(y))
    case uint32:
    return o.uin(uint64(x), uint64(y))
    case uint64:
    return o.uin(uint64(x), uint64(y))
    case uint:
    return o.uin(uint64(x), uint64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case uint:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.uin(uint64(x), timeToUint(y))
    case byte:
    return o.uin(uint64(x), uint64(y))
    case uint16:
    return o.uin(uint64(x), uint64(y))
    case uint32:
    return o.uin(uint64(x), uint64(y))
    case uint64:
    return o.uin(uint64(x), uint64(y))
    case uint:
    return o.uin(uint64(x), uint64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case int8:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.sin(int64(x), timeToInt(y))
    case byte:
    return o.sin(int64(x), int64(y))
    case uint16:
    return o.sin(int64(x), int64(y))
    case uint32:
    return o.sin(int64(x), int64(y))
    case uint64:
    return o.sin(int64(x), int64(y))
    case uint:
    return o.sin(int64(x), int64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case int16:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.sin(int64(x), timeToInt(y))
    case byte:
    return o.sin(int64(x), int64(y))
    case uint16:
    return o.sin(int64(x), int64(y))
    case uint32:
    return o.sin(int64(x), int64(y))
    case uint64:
    return o.sin(int64(x), int64(y))
    case uint:
    return o.sin(int64(x), int64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case int32:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.sin(int64(x), timeToInt(y))
    case byte:
    return o.sin(int64(x), int64(y))
    case uint16:
    return o.sin(int64(x), int64(y))
    case uint32:
    return o.sin(int64(x), int64(y))
    case uint64:
    return o.sin(int64(x), int64(y))
    case uint:
    return o.sin(int64(x), int64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case int64:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.sin(int64(x), timeToInt(y))
    case byte:
    return o.sin(int64(x), int64(y))
    case uint16:
    return o.sin(int64(x), int64(y))
    case uint32:
    return o.sin(int64(x), int64(y))
    case uint64:
    return o.sin(int64(x), int64(y))
    case uint:
    return o.sin(int64(x), int64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case int:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.sin(int64(x), timeToInt(y))
    case byte:
    return o.sin(int64(x), int64(y))
    case uint16:
    return o.sin(int64(x), int64(y))
    case uint32:
    return o.sin(int64(x), int64(y))
    case uint64:
    return o.sin(int64(x), int64(y))
    case uint:
    return o.sin(int64(x), int64(y))
    case int8:
    return o.sin(int64(x), int64(y))
    case int16:
    return o.sin(int64(x), int64(y))
    case int32:
    return o.sin(int64(x), int64(y))
    case int64:
    return o.sin(int64(x), int64(y))
    case int:
    return o.sin(int64(x), int64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case float32:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.fl(float64(x), timeToFloat(y))
    case byte:
    return o.fl(float64(x), float64(y))
    case uint16:
    return o.fl(float64(x), float64(y))
    case uint32:
    return o.fl(float64(x), float64(y))
    case uint64:
    return o.fl(float64(x), float64(y))
    case uint:
    return o.fl(float64(x), float64(y))
    case int8:
    return o.fl(float64(x), float64(y))
    case int16:
    return o.fl(float64(x), float64(y))
    case int32:
    return o.fl(float64(x), float64(y))
    case int64:
    return o.fl(float64(x), float64(y))
    case int:
    return o.fl(float64(x), float64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  case float64:
  switch y := b.(type) {
    case nil:
    return o.nl(x, y)
    case bool:
    if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt
    case string:
    nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)
    case time.Time:
    return o.fl(float64(x), timeToFloat(y))
    case byte:
    return o.fl(float64(x), float64(y))
    case uint16:
    return o.fl(float64(x), float64(y))
    case uint32:
    return o.fl(float64(x), float64(y))
    case uint64:
    return o.fl(float64(x), float64(y))
    case uint:
    return o.fl(float64(x), float64(y))
    case int8:
    return o.fl(float64(x), float64(y))
    case int16:
    return o.fl(float64(x), float64(y))
    case int32:
    return o.fl(float64(x), float64(y))
    case int64:
    return o.fl(float64(x), float64(y))
    case int:
    return o.fl(float64(x), float64(y))
    case float32:
    return o.fl(float64(x), float64(y))
    case float64:
    return o.fl(float64(x), float64(y))
    default:
      return o.dflt
  }
  default:
    return o.dflt
}
}
}