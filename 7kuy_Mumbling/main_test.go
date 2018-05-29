package main

import (
	"testing"
)

type testCase struct {
	str    string
	result string
}

var testCases = []testCase{
	{"abcd", "A-Bb-Ccc-Dddd"},
	{"RqaEzty", "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"},
	{"cwAt", "C-Ww-Aaa-Tttt"},
}

func TestAccum(t *testing.T) {
	for _, testCase := range testCases {
		if Accum(testCase.str) != testCase.result {
			t.Errorf("%s should be %s\n", testCase.str, testCase.result)
		}
	}
}
