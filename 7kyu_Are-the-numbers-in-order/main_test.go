package main

import (
	"reflect"
	"runtime"
	"sort"
	"testing"
)

var testCases = [][]int{
	[]int{1, 2, 4, 7, 19},
	[]int{1, 2, 3, 4, 5},
	[]int{1, 6, 10, 18, 2, 4, 20},
	[]int{9, 8, 7, 6, 5, 4, 3, 2, 1},
	[]int{-100, -99, -0, -89, -35, -22, -19, -16, -5, -1, 0, 1, 2, 3, 4, 5, 6, 7, 10, 13, 15, 91, 100},
}

func TestInAscOrder(t *testing.T) {
	solutions := []interface{}{InAscOrder1, InAscOrder2}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			if solution.(func([]int) bool)(testCase) != sort.IntsAreSorted(testCase) {
				t.Errorf("%s %v should be %v\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase, sort.IntsAreSorted(testCase))
			}
		}
	}
}

func BenchmarkInAscOrder1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range testCases {
			InAscOrder1(v)
		}
	}
}

func BenchmarkInAscOrder2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range testCases {
			InAscOrder2(v)
		}
	}
}
