// You are given an array strarr of strings and an integer k. Your task is to return the first longest string consisting of k consecutive strings taken in the array.

// #Example: longest_consec(["zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"], 2) --> "abigailtheta"

// n being the length of the string array, if n = 0 or k > n or k <= 0 return "".

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

	fmt.Println(LongestConsec1([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	fmt.Println(LongestConsec2([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	fmt.Println(LongestConsec3([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
	fmt.Println(LongestConsec4([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))

}

func LongestConsec1(strarr []string, k int) (result string) {
	var curStr string
	for i := 0; i < len(strarr)+1-k; i++ {
		for j := i; j < i+k; j++ {
			curStr += strarr[j]
		}

		if len(curStr) > len(result) {
			result = curStr
		}
		curStr = ""
	}

	return
}

// This is the fastest function
func LongestConsec2(strarr []string, k int) (result string) {
	for i := 0; i < len(strarr)+1-k; i++ {
		curStr := strings.Join(strarr[i:i+k], "")
		if len(curStr) > len(result) {
			result = curStr
		}
	}
	return
}

func LongestConsec3(strarr []string, k int) (result string) {
	var bb bytes.Buffer // A Buffer needs no initialization.
	for i := 0; i < len(strarr)+1-k; i++ {
		for j := i; j < i+k; j++ {
			bb.WriteString(strarr[j])
		}

		if bb.Len() > len(result) {
			result = bb.String()
		}
		bb.Reset()
	}

	return
}

func LongestConsec4(strarr []string, k int) (result string) {
	var bb bytes.Buffer
	for i := 0; i < len(strarr)+1-k; i++ {
		bb.WriteString(strings.Join(strarr[i:i+k], ""))
		if bb.Len() > len(result) {
			result = bb.String()
		}
		bb.Reset()
	}
	return
}
