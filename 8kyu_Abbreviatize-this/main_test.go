package main

import "testing"

type testCase struct {
	Word         string
	Abbreviation string
}

var testCases = []testCase{
	{"Globalization", "G11n"},
	{"Internationalization", "I18n"},
	{"Scherpenhuizen", "S12n"},
	{"Four", "F2r"},
	{"The", "The"},
	{"A", "A"},
}

func TestAbbreviatize1(t *testing.T) {
	for _, testCase := range testCases {
		if Abbreviatize1(testCase.Word) != testCase.Abbreviation {
			t.Errorf("%q should be %q for the %q\n", Abbreviatize1(testCase.Word), testCase.Abbreviation, testCase.Word)
		}
	}
}
func TestAbbreviatize2(t *testing.T) {
	for _, testCase := range testCases {
		if Abbreviatize1(testCase.Word) != testCase.Abbreviation {
			t.Errorf("%q should be %q for the %q\n", Abbreviatize1(testCase.Word), testCase.Abbreviation, testCase.Word)
		}
	}
}
func TestAbbreviatize3(t *testing.T) {
	for _, testCase := range testCases {
		if Abbreviatize1(testCase.Word) != testCase.Abbreviation {
			t.Errorf("%q should be %q for the %q\n", Abbreviatize1(testCase.Word), testCase.Abbreviation, testCase.Word)
		}
	}
}
func TestAbbreviatize4(t *testing.T) {
	for _, testCase := range testCases {
		if Abbreviatize1(testCase.Word) != testCase.Abbreviation {
			t.Errorf("%q should be %q for the %q\n", Abbreviatize1(testCase.Word), testCase.Abbreviation, testCase.Word)
		}
	}
}

func BenchmarkAbbreviatize1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Abbreviatize1(testCase.Word)
		}
	}
}

func BenchmarkAbbreviatize2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Abbreviatize1(testCase.Word)
		}
	}
}

func BenchmarkAbbreviatize3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Abbreviatize1(testCase.Word)
		}
	}
}

func BenchmarkAbbreviatize4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Abbreviatize1(testCase.Word)
		}
	}
}
