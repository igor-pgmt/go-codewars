package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	number int
	answer int
}

var testCases = []testCase{
	{1234567, 87},
	{8529, 79},
	{85299258, 31},
	{5634, 57},
}

func TestThirt(t *testing.T) {
	solutions := []interface{}{Thirt1, Thirt2}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			result := solution.(func(int) int)(testCase.number)
			if result != testCase.answer {
				t.Errorf("%s \n Argument:%d Answer: %d Should be: %d\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase.number, result, testCase.answer)
			}
		}
	}
}

func BenchmarkThirt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			Thirt1(v.number)
		}
	}
}

func BenchmarkThirt2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range testCases {
			Thirt2(v.number)
		}
	}
}
