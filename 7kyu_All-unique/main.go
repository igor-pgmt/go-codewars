// https://www.codewars.com/kata/553e8b195b853c6db4000048

// Write a program to determine if a string contains all unique characters.
// Return true if it does and false otherwise.

// The string may contain any of the 128 ASCII characters.

package main

import "fmt"

func main() {
	fmt.Println(CheckUnique1("ab")) // true
	fmt.Println(CheckUnique1("aa")) // false
	fmt.Println(CheckUnique2("ab")) // true
	fmt.Println(CheckUnique2("aa")) // false
	fmt.Println(CheckUnique3("ab")) // true
	fmt.Println(CheckUnique3("aa")) // false
}

// This is the fastest function
func CheckUnique1(str string) bool {
	list := map[rune]int{}
	for _, val := range str {
		if list[val] == 0 {
			list[val]++
		} else {
			return false
		}
	}
	return true
}

func CheckUnique2(str string) bool {
	list := map[rune]int{}
	for _, val := range str {
		list[val]++
	}
	for _, val := range list {
		if val > 1 {
			return false
		}
	}
	return true
}

func CheckUnique3(str string) bool {
	var m = map[rune]bool{}
	for _, w := range str {
		m[w] = true
	}
	return len(m) == len(str)
}
