package main

import (
	"testing"
)

type testCase struct {
	items  []interface{}
	k      int
	result []interface{}
}

var testCases = []testCase{
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 1, []interface{}{1, 2, 3, 4, 5, 6, 7}},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 2, []interface{}{2, 4, 6, 1, 5, 3, 7}},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 3, []interface{}{3, 6, 2, 7, 5, 1, 4}},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 4, []interface{}{4, 1, 6, 5, 7, 3, 2}},
	{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 10, []interface{}{3, 7, 6, 2, 4, 1, 5}},
}

func TestJosephus(t *testing.T) {
	for _, testCase := range testCases {
		items := make([]interface{}, len(testCase.items))
		copy(items, testCase.items)
		result := Josephus(items, testCase.k)
		if !compareSlices(result, testCase.result) {
			t.Errorf("\nArguments:%v, %d Answer: %v Should be: %v\n", testCase.items, testCase.k, result, testCase.result)
		}
	}
}

func compareSlices(a, b []interface{}) bool {
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
