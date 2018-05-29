package main

import (
	"reflect"
	"runtime"
	"testing"
)

type testCase struct {
	str      string
	isUnique bool
}

var testCases = []testCase{
	{"asdfghkjklwpeeoriuddnvxcvd", false},
	{"abcdefghijklmnopqrstuvwxyz", true},
	{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", true},
	{"asdfghkjklwpeeoriuddnvxcvdFASDFJKASDFHKJASFHKSADFKJA", false},
	{"123", true},
	{"122", false},
}

func TestCheckUnique(t *testing.T) {
	solutions := []interface{}{CheckUnique1, CheckUnique2, CheckUnique3}
	for _, solution := range solutions {
		for _, testCase := range testCases {
			if solution.(func(string) bool)(testCase.str) != testCase.isUnique {
				t.Errorf("%s %s should be %v\n", runtime.FuncForPC(reflect.ValueOf(solution).Pointer()).Name(), testCase.str, testCase.isUnique)
			}
		}
	}
}

func BenchmarkCheckUnique1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range testCases {
			CheckUnique1(v.str)
		}
	}
}
func BenchmarkCheckUnique2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range testCases {
			CheckUnique2(v.str)
		}
	}
}
func BenchmarkCheckUnique3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, v := range testCases {
			CheckUnique3(v.str)
		}
	}
}
