package main

import (
	"math"
	"reflect"
	"runtime"
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	solutions := []interface{}{EvenOrOdd1, EvenOrOdd2}
	for _, solution := range solutions {
		for i := -1000; i < 1000; i++ {
			s := solution.(func(int) string)(i)
			if s != []string{"Even", "Odd"}[int(math.Abs(float64(i%2)))] {
				t.Errorf("%s %d != %s\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), i, s)
			}
		}
	}
}

func BenchmarkEvenOrOdd1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			EvenOrOdd1(i)
		}
	}
}
func BenchmarkEvenOrOdd2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000; i++ {
			EvenOrOdd2(i)
		}
	}
}
