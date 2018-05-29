package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	str      string
	reversed string
}

var testCases = []testCase{
	{"", ""},
	{"0", "0"},
	{"qw", "wq"},
	{"cdE", "Edc"},
	{"8irT", "Tri8"},
	{"1w2e3", "3e2w1"},
	{"QWErty", "ytrEWQ"},
	{"123456pzl", "lzp654321"},
}

func TestSolutions(t *testing.T) {
	solutions := []interface{}{Solution1, Solution2, Solution3, Solution4, Solution5}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			if solution.(func(string) string)(testCase.str) != testCase.reversed {
				t.Errorf("%s %s != %s\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase.str, testCase.reversed)
			}
		}
	}
}

func BenchmarkSolution1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Solution1(testCase.str)
		}
	}
}
func BenchmarkSolution2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Solution2(testCase.str)
		}
	}
}
func BenchmarkSolution3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Solution3(testCase.str)
		}
	}
}
func BenchmarkSolution4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Solution4(testCase.str)
		}
	}
}
func BenchmarkSolution5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Solution5(testCase.str)
		}
	}
}
