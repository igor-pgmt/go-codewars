//https://www.codewars.com/kata/5667e8f4e3f572a8f2000039

// This time no story, no theory. The examples below show you how to write function accum:

// Examples:

// accum("abcd") // -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") // -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") // -> "C-Ww-Aaa-Tttt"

// The parameter of accum is a string which includes only letters from a..z and A..Z.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Accum("abc")) // A-Bb-Ccc
}

func Accum(s string) string {
	s = strings.ToLower(s)
	resultSlc := []string{}
	for i, letter := range s {
		resultSlc = append(resultSlc, strings.Title(strings.Repeat(string(letter), i+1)))
	}
	return strings.Join(resultSlc, "-")
}
