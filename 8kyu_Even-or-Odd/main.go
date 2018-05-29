// https://www.codewars.com/kata/even-or-odd/

// Create a function that takes an integer as an argument
// and returns "Even" for even numbers or "Odd" for odd numbers.

package main

import "fmt"

func main() {

	fmt.Println(EvenOrOdd1(1))
	fmt.Println(EvenOrOdd1(2))
	fmt.Println(EvenOrOdd2(1))
	fmt.Println(EvenOrOdd2(2))

}

// This is the fastest function
func EvenOrOdd1(number int) string {
	if number%2 == 0 {
		return "Even"
	}
	return "Odd"
}

func EvenOrOdd2(number int) string {
	return []string{"Even", "Odd"}[number&1]
}
