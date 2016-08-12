// This program generates the type switch statements needed by goexpr to
// implement type casting
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

var (
	uints      = []string{"byte", "uint16", "uint32", "uint64", "uint"}
	sints      = []string{"int8", "int16", "int32", "int64", "int"}
	ints       = append(uints, sints...)
	floats     = []string{"float32", "float64"}
	numbers    = append(ints, floats...)
	otherTypes = []string{"nil", "bool", "string", "time.Time"}
	types      = append(otherTypes, numbers...)
)

func main() {
	var buf bytes.Buffer
	buf.WriteString(`package goexpr

import (
  "time"
)

func buildOp(_o string) func(interface{}, interface{}) interface{} {
	o := ops[_o]
	return func(a interface{}, b interface{}) interface{} {
`)
	buf.WriteString("switch x := a.(type) {\n")
	for _, t1 := range types {
		buf.WriteString(fmt.Sprintf("  case %v:\n", t1))
		buf.WriteString("  switch y := b.(type) {\n")
		for _, t2 := range types {
			buf.WriteString(fmt.Sprintf("    case %v:\n", t2))
			ex := exprFor(t1, t2)
			buf.WriteString(fmt.Sprintf("    %v\n", ex))
		}
		buf.WriteString("    default:\n")
		buf.WriteString("      return o.dflt\n")
		buf.WriteString("  }\n")
	}
	buf.WriteString("  default:\n")
	buf.WriteString("    return o.dflt\n")
	buf.WriteString("}\n}\n}")
	outfile := filepath.Join("../ops.go")
	err := ioutil.WriteFile(outfile, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote to %v\n", outfile)
}

func exprFor(t1 string, t2 string) string {
	switch t1 {
	case "nil":
		switch t2 {
		case "nil":
			return `return o.nl(x, y)`
		case "bool":
			return `return o.bl(false, y)`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `return o.uin(0, uint64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `return o.sin(0, int64(y))`
		case "float32", "float64":
			return `return o.fl(0, float64(y))`
		case "string":
			return `return o.st("", y)`
		case "time.Time":
			return `return o.ts(zeroTime, y)`
		default:
			return `return o.dflt`
		}
	case "bool":
		switch t2 {
		case "nil":
			return `return o.bl(x, false)`
		case "bool":
			return `return o.bl(x, y)`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return `if y == 1 {
  return o.bl(x, true)
}
if y == 0 {
	return o.bl(x, false)
}
return o.dflt`
		case "string":
			return `bv, ok := strToBool(y)
if !ok {
	return o.dflt
}
return o.bl(x, bv)`
		case "time.Time":
			return `return o.bl(x, timeToBool(y))`
		default:
			return `return o.dflt`
		}
	case "byte", "uint16", "uint32", "uint64", "uint":
		switch t2 {
		case "nil":
			return `return o.nl(x, y)`
		case "bool":
			return `if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `return o.uin(uint64(x), uint64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `return o.sin(int64(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)`
		case "time.Time":
			return `return o.uin(uint64(x), timeToUint(y))`
		default:
			return `return o.dflt`
		}
	case "int8", "int16", "int32", "int64", "int":
		switch t2 {
		case "nil":
			return `return o.nl(x, y)`
		case "bool":
			return `if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int":
			return `return o.sin(int64(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)`
		case "time.Time":
			return `return o.sin(int64(x), timeToInt(y))`
		default:
			return `return o.dflt`
		}
	case "float32", "float64":
		switch t2 {
		case "nil":
			return `return o.nl(x, y)`
		case "bool":
			return `if x == 1 {
  return o.bl(true, y)
}
if x == 0 {
	return o.bl(false, y)
}
return o.dflt`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `nv, ok := strToFloat(y)
if !ok {
	return o.dflt
}
return o.fl(float64(x), nv)`
		case "time.Time":
			return `return o.fl(float64(x), timeToFloat(y))`
		default:
			return `return o.dflt`
		}
	case "string":
		switch t2 {
		case "nil":
			return `return o.st(x, "")`
		case "bool":
			return `bv, ok := strToBool(x)
if !ok {
	return o.dflt
}
return o.bl(bv, y)`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))`
		case "float32", "float64":
			return `nv, ok := strToFloat(x)
if !ok {
	return o.dflt
}
return o.fl(nv, float64(y))`
		case "string":
			return `return o.st(x, y)`
		case "time.Time":
			return `return o.st(x, y.String())`
		default:
			return `return o.dflt`
		}
	case "time.Time":
		switch t2 {
		case "nil":
			return `return o.ts(x, zeroTime)`
		case "bool":
			return `return o.bl(timeToBool(x), y)`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `return o.uin(timeToUint(x), uint64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `return o.sin(timeToInt(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(timeToFloat(x), float64(y))`
		case "string":
			return `return o.st(x.String(), y)`
		case "time.Time":
			return `return o.ts(x, y)`
		default:
			return `return o.dflt`
		}
	}
	return `return o.dflt`
}
