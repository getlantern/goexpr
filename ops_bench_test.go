package goexpr

import (
	"reflect"
	"testing"
)

func BenchmarkTypeSwitch(b *testing.B) {
	var x interface{} = int16(5)
	var y interface{} = int16(6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		switch v := x.(type) {
		case int16:
			switch v2 := y.(type) {
			case int16:
				if v == v2 {
				}
			}
		}
	}
}

func BenchmarkCast(b *testing.B) {
	var x interface{} = int16(5)
	var y interface{} = int16(6)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v int64
		var v2 int64
		switch _v := x.(type) {
		case int16:
			v = int64(_v)
		}
		switch _v2 := y.(type) {
		case int16:
			v2 = int64(_v2)
		}
		if v == v2 {
		}
	}
}

func BenchmarkReflect(b *testing.B) {
	var x interface{} = int16(5)
	var y interface{} = int16(6)

	int16Type := reflect.TypeOf(int16(0))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var v int64
		var v2 int64
		if reflect.TypeOf(x) == int16Type {
			v = int64(x.(int16))
		}
		if reflect.TypeOf(y) == int16Type {
			v2 = int64(y.(int16))
		}
		if v == v2 {
		}
	}
}
