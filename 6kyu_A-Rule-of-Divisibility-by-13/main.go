// https://www.codewars.com/kata/a-rule-of-divisibility-by-13/go

package main

import (
	"fmt"
	"strconv"
)

var division = []int{1, 10, 9, 12, 3, 4}

func main() {
	fmt.Println(Thirt1(123)) // 32
	fmt.Println(Thirt2(123)) // 32
}

// This is the fastest function
func Thirt1(n int) int {
	sum := 0
	number := n

	for i := 1; number > 0; i *= 10 {
		sum += (i % 13) * (number % 10)
		number /= 10
	}

	if n == sum {
		return n
	}

	return Thirt1(sum)
}

func Thirt2(n int) int {
	strings := strconv.Itoa(n)
	var result, buffer int

More:
	result = buffer
	buffer = 0

	for i, j := len(strings)-1, 0; i >= 0; i, j = i-1, j+1 {
		nyint, _ := strconv.Atoi(string(strings[i]))
		buffer += nyint * division[j%6]
	}

	if result != buffer {
		strings = strconv.Itoa(buffer)
		goto More
	}
	return result
}
