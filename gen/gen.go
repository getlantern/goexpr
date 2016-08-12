// This program generates the type switch statements needed by goexpr to
// implement type casting
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"
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
	"strings"
  "time"
)

var _ops = map[string]func(interface{}, interface{}) interface{} {
`)
	for _, _o := range []string{"==", "!=", "<", "<=", ">", ">=", "+", "-", "*", "/"} {
		o := ops[_o]
		buf.WriteString(fmt.Sprintf(`"%v": func(a interface{}, b interface{}) interface{} {
`, _o))
		buf.WriteString("switch v1 := a.(type) {\n")
		for _, t1 := range types {
			buf.WriteString(fmt.Sprintf("  case %v:\n", t1))
			buf.WriteString("  switch v2 := b.(type) {\n")
			for _, t2 := range types {
				buf.WriteString(fmt.Sprintf("    case %v:\n", t2))
				if t1 == t2 {
					opString := opForType(o, t1)
					if opString == "" {
						buf.WriteString(fmt.Sprintf("    return %v\n", o.dflt))
						continue
					}
					buf.WriteString(fmt.Sprintf("    x, y := v1, v2\n"))
					buf.WriteString(fmt.Sprintf("    return %v\n", opString))
					continue
				}
				_t1 := t1
				_t2 := t2
				v1, v2 := "v1", "v2"
				x, y := "x", "y"
				commonType, expr1, expr2 := params(_t1, _t2)
				if commonType == "" {
					commonType, expr1, expr2 = params(_t2, _t1)
					if commonType == "" {
						panic(fmt.Errorf("Unable to find type conversion for %v -> %v", t1, t2))
					}
					v1, v2 = v2, v1
					x, y = y, x
				}
				opString := opForType(o, commonType)
				if opString == "" {
					buf.WriteString(fmt.Sprintf("    return %v\n", o.dflt))
					continue
				}
				buf.WriteString(fmt.Sprintf(`var %v %v
var xok bool
%v, xok = %v
var %v %v
var yok bool
%v, yok = %v
if !xok || !yok {
  return %v
}
return %v
`, x, commonType, x, render(expr1, v1, v2), y, commonType, y, render(expr2, v1, v2), defaultFor(commonType), opString))
			}
			buf.WriteString("    default:\n")
			buf.WriteString(fmt.Sprintf("      return %v\n", o.dflt))
			buf.WriteString("  }\n")
		}
		buf.WriteString("  default:\n")
		buf.WriteString(fmt.Sprintf("    return %v\n", o.dflt))
		buf.WriteString("}\n},\n")
	}
	buf.WriteString("}\n")
	outfile := filepath.Join("../ops.go")
	err := ioutil.WriteFile(outfile, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote to %v\n", outfile)
}

func opForType(o op, t string) string {
	switch t {
	case "nil":
		return o.nl
	case "bool":
		return o.bl
	case "string":
		return o.st
	case "time.Time":
		return o.ts
	default:
		return o.nu
	}
}

func defaultFor(t string) string {
	switch t {
	case "bool":
		return `false`
	case "string":
		return `""`
	case "time.Time":
		return `zeroTime`
	default:
		return `0`
	}
}

func render(tmpl, self, other string) string {
	buf := bytes.NewBuffer(nil)
	err := template.Must(template.New("tmpl").Parse(tmpl)).Execute(buf, map[string]interface{}{"s": self, "o": other})
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func params(from string, to string) (commonType string, expr1 string, expr2 string) {
	switch from {
	case "nil":
		switch to {
		case "nil":
			return to, `{{.s}}, true`, `{{.o}}, true`
		case "bool":
			return to, `false, true`, `{{.o}}, true`
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return to, `0, true`, `{{.o}}, true`
		case "string":
			return to, `"", true`, `{{.o}}, true`
		case "time.Time":
			return to, `time.Time{}, true`, `{{.o}}, true`
		}
	case "bool":
		switch to {
		case "byte", "uint16", "uint32", "uint64", "uint", "int8", "int16", "int32", "int64", "int", "float32", "float64":
			return "bool", `{{.s}}, true`, `{{.o}} == 1, true`
		case "string":
			return "bool", `{{.s}}, true`, `strToBool({{.o}})`
		case "time.Time":
			return "bool", `{{.s}}, true`, `timeToBool({{.o}}), true`
		}
	case "string":
		switch to {
		case "byte", "uint16", "uint32", "uint64", "uint":
			return "uint64", `strToUint({{.s}})`, `uint64({{.o}}), true`
		case "int8", "int16", "int32", "int64", "int":
			return "int64", `strToInt({{.s}})`, `int64({{.o}}), true`
		case "float32", "float64":
			return "float64", `strToFloat({{.s}})`, `float64({{.o}}), true`
		case "time.Time":
			return "string", `{{.s}}, true`, `{{.o}}.String(), true`
		}
	case "byte", "uint16", "uint32", "uint64", "uint":
		switch to {
		case "byte", "uint16", "uint32", "uint64", "uint":
			return "uint64", `uint64({{.s}}), true`, `uint64({{.o}}), true`
		case "int8", "int16", "int32", "int64", "int":
			return "int64", `int64({{.s}}), true`, `int64({{.o}}), true`
		case "float32", "float64":
			return "float64", `float64({{.s}}), true`, `float64({{.o}}), true`
		case "time.Time":
			return "time.Time", `time.Unix(int64({{.s}}), 0), true`, `{{.o}}, true`
		}
	case "int8", "int16", "int32", "int64", "int":
		switch to {
		case "int8", "int16", "int32", "int64", "int":
			return "int64", `int64({{.s}}), true`, `int64({{.o}}), true`
		case "float32", "float64":
			return "float64", `float64({{.s}}), true`, `float64({{.o}}), true`
		case "time.Time":
			return "time.Time", `time.Unix(int64({{.s}}), 0), true`, `{{.o}}, true`
		}
	case "float32", "float64":
		switch to {
		case "float32", "float64":
			return "float64", `float64({{.s}}), true`, `float64({{.o}}), true`
		case "time.Time":
			return "time.Time", `time.Unix(int64({{.s}}), 0), true`, `{{.o}}, true`
		}
	}

	return "", "", ""
}

type op struct {
	nl   string
	nu   string
	bl   string
	st   string
	ts   string
	dflt string
}

var ops = map[string]op{
	"==": op{
		nl:   "true",
		nu:   "x == y",
		bl:   "x == y",
		st:   "x == y",
		ts:   "x == y",
		dflt: "false",
	},
	"!=": op{
		nl:   "false",
		nu:   "x != y",
		bl:   "x != y",
		st:   "x == y",
		ts:   "x != y",
		dflt: "false",
	},
	"<": op{
		nl:   "false",
		nu:   "x < y",
		bl:   "!x && y",
		st:   "strings.Compare(x, y) < 0",
		ts:   "x.Before(y)",
		dflt: "false",
	},
	"<=": op{
		nl:   "true",
		nu:   "x <= y",
		bl:   "true",
		st:   "strings.Compare(x, y) <= 0",
		ts:   "!x.After(y)",
		dflt: "false",
	},
	">": op{
		nl:   "false",
		nu:   "x > y",
		bl:   "x && !y",
		st:   "strings.Compare(x, y) > 0",
		ts:   "x.After(y)",
		dflt: "false",
	},
	">=": op{
		nl:   "true",
		nu:   "x >= y",
		bl:   "true",
		st:   "strings.Compare(x, y) >= 0",
		ts:   "!x.Before(y)",
		dflt: "false",
	},
	"+": op{
		nu:   "x + y",
		st:   "x + y",
		dflt: "0",
	},
	"-": op{
		nu:   "x - y",
		dflt: "0",
	},
	"*": op{
		nu:   "x * y",
		dflt: "0",
	},
	"/": op{
		nu:   "div(x, y)",
		dflt: "0",
	},
}
