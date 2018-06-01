package main

import (
	"testing"
)

type testCase struct {
	input, answer []int
}

var testCases = []testCase{
	{[]int{22, -6, 32, 82, 9, 25}, []int{-6, 32, 25}},
	{[]int{68, -1, 1, -7, 10, 10}, []int{-1, 10}},
	{[]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68}, []int{-85, 72, 0, 68}},
}

func TestMultipleOfIndex(t *testing.T) {
	for _, testCase := range testCases {
		result := MultipleOfIndex(testCase.input)
		answer := testCase.answer
		if !compareSlices(result, answer) {
			t.Errorf("%v should be %v\n", result, answer)
		}
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
