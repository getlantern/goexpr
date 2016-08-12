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

func applyOp(o op, a interface{}, b interface{}) interface{} {
`)
	buf.WriteString("switch x := a.(type) {\n")
	for _, t1 := range types {
		buf.WriteString(fmt.Sprintf("  case %v:\n", t1))
		if t1 == "nil" {
			buf.WriteString("    return o.nl(a, b)\n")
			continue
		}
		buf.WriteString("  switch y := b.(type) {\n")
		for _, t2 := range types {
			ex := exprFor(t1, t2)
			if ex != "" {
				buf.WriteString(fmt.Sprintf("    case %v:\n", t2))
				buf.WriteString(fmt.Sprintf("    %v\n", ex))
			}
		}
		buf.WriteString("  }\n")
	}
	buf.WriteString(`}
return o.dflt
}`)
	outfile := filepath.Join("../apply.go")
	err := ioutil.WriteFile(outfile, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote to %v\n", outfile)
}

func exprFor(t1 string, t2 string) string {
	switch t1 {
	case "nil":
		return `return o.nl(a, b)`
	case "bool":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "bool":
			return `return o.bl(x, y)`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `return o.uin(boolToUInt(x), uint64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `return o.sin(boolToInt(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(boolToFloat(x), float64(y))`
		case "string":
			return `yb, ok := strToBool(y)
if ok {
	return o.bl(x, yb)
}`
		}
	case "byte", "uint16", "uint32", "uint64", "uint":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "bool":
			return `return o.uin(uint64(x), boolToUInt(y))`
		case "byte", "uint16", "uint32", "uint64", "uint":
			return `return o.uin(uint64(x), uint64(y))`
		case "int8", "int16", "int32", "int64", "int":
			return `return o.sin(int64(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `yf, ok := strToFloat(y)
if ok {
	return o.fl(float64(x), yf)
}`
		}
	case "int8", "int16", "int32", "int64", "int":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "bool":
			return `return o.sin(int64(x), boolToInt(y))`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int":
			return `return o.sin(int64(x), int64(y))`
		case "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `yf, ok := strToFloat(y)
if ok {
	return o.fl(float64(x), yf)
}`
		}
	case "float32", "float64":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "bool":
			return `return o.fl(float64(x), boolToFloat(y))`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return `return o.fl(float64(x), float64(y))`
		case "string":
			return `yf, ok := strToFloat(y)
if ok {
	return o.fl(float64(x), yf)
}`
		}
	case "string":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "bool":
			return `xb, ok := strToBool(x)
if ok {
	return o.bl(xb, y)
}`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return `xf, ok := strToFloat(x)
if ok {
	return o.fl(xf, float64(y))
}`
		case "string":
			return `return o.st(x, y)`
		case "time.Time":
			return `return o.st(x, y.String())`
		}
	case "time.Time":
		switch t2 {
		case "nil":
			return `return o.nl(a, b)`
		case "string":
			return `return o.st(x.String(), y)`
		case "time.Time":
			return `return o.ts(x, y)`
		}
	}

	return ""
}
