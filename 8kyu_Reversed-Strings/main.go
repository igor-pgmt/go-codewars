// https://www.codewars.com/kata/reversed-strings

// Complete the solution so that it reverses the string value passed into it.

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// This is the fastest function
func Solution1(word string) string {
	runes := []rune(word)
	sLen := len(runes)
	for i := 0; i < sLen/2; i++ {
		runes[i], runes[sLen-1-i] = runes[sLen-1-i], runes[i]
	}
	return string(runes)
}

func Solution2(word string) string {
	cs := make([]rune, utf8.RuneCountInString(word))
	i := len(cs)
	for _, c := range word {
		i--
		cs[i] = c
	}
	return string(cs)
}

func Solution3(word string) string {
	rr := []rune(word)
	l := len(rr)
	for i, j := 0, l-1; i < l/2; i, j = i+1, j-1 {
		rr[i], rr[j] = rr[j], rr[i]
	}
	return string(rr)
}

func Solution4(word string) string {
	var reversed []string
	for i := range word {
		reversed = append(reversed, string(word[len(word)-1-i]))
	}
	return strings.Join(reversed, "")
}

func Solution5(word string) (result string) {
	for _, c := range word {
		result = string(c) + result
	}
	return result
}

func main() {
	word := "word"
	fmt.Println(Solution1(word))
	fmt.Println(Solution2(word))
	fmt.Println(Solution3(word))
	fmt.Println(Solution4(word))
	fmt.Println(Solution5(word))
}
