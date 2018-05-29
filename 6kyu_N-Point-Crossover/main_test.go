package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	ns   []int
	xs   []int
	ys   []int
	xans []int
	yans []int
}

var testCases = []testCase{
	{[]int{}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}},
	{[]int{1}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}, []int{1, 2, 2, 2, 2}, []int{2, 1, 1, 1, 1}},
	{[]int{1, 1}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}, []int{1, 2, 2, 2, 2}, []int{2, 1, 1, 1, 1}},
	{[]int{1, 3}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2}, []int{1, 2, 2, 1, 1}, []int{2, 1, 1, 2, 2}},
	{[]int{1, 3, 5}, []int{1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2}, []int{1, 2, 2, 1, 1, 2, 2}, []int{2, 1, 1, 2, 2, 1, 1}},
	{[]int{3, 5, 1, 1, 3}, []int{1, 1, 1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2, 2, 2}, []int{1, 2, 2, 1, 1, 2, 2}, []int{2, 1, 1, 2, 2, 1, 1}},
}

func TestCheckCrossover(t *testing.T) {
	solutions := []interface{}{Crossover1, Crossover2}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			xs := make([]int, len(testCase.xs))
			ys := make([]int, len(testCase.ys))
			copy(xs, testCase.xs)
			copy(ys, testCase.ys)
			xans, yans := solution.(func([]int, []int, []int) ([]int, []int))(testCase.ns, xs, ys)
			if !compareSlices(xans, testCase.xans) || !compareSlices(yans, testCase.yans) {
				t.Errorf("%s %v and %v should be %v and %v\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), xs, ys, testCase.xans, testCase.yans)
			}
		}
	}
}

func BenchmarkCrossover1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			Crossover1(v.ns, v.xs, v.ys)
		}
	}
}

func BenchmarkCrossover2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			Crossover2(v.ns, v.xs, v.ys)
		}
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
