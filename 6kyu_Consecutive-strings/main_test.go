package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	Strings []string
	Length  int
	Result  string
}

var testCases = []testCase{
	{[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2, "abigailtheta"},
	{[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}, 1,
		"oocccffuucccjjjkkkjyyyeehh"},
	{[]string{}, 3, ""},
	{[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"}, 2,
		"wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck"},
}

func TestLongestConsec(t *testing.T) {
	solutions := []interface{}{LongestConsec1, LongestConsec2, LongestConsec3, LongestConsec4}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			result := solution.(func([]string, int) string)(testCase.Strings, testCase.Length)
			if result != testCase.Result {
				t.Errorf("%s\n Arguments: %v, %d Result:%s Should be: %s\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase.Strings, testCase.Length, result, testCase.Result)
			}
		}
	}
}

func BenchmarkLongestConsec1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			LongestConsec1(testCase.Strings, testCase.Length)
		}
	}
}
func BenchmarkLongestConsec2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			LongestConsec2(testCase.Strings, testCase.Length)
		}
	}
}
func BenchmarkLongestConsec3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			LongestConsec3(testCase.Strings, testCase.Length)
		}
	}
}
func BenchmarkLongestConsec4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			LongestConsec4(testCase.Strings, testCase.Length)
		}
	}
}
