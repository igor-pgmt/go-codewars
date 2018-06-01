// https://www.codewars.com/kata/511f11d355fe575d2c000001
// Two Oldest Ages
// The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].
// The order of the numbers passed in could be any order. The array will always include at least 2 items.
// For example:
// TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}

package main

import "fmt"

func main() {

	fmt.Println(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}))

}

func TwoOldestAges(ages []int) (result [2]int) {
	for _, age := range ages {
		switch {
		case age > result[1]:
			result[0], result[1] = result[1], age
		case age > result[0]:
			result[0] = age
		}
	}
	return
}
