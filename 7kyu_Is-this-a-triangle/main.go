// https://www.codewars.com/kata/56606694ec01347ce800001b
// Is this a triangle?
// Implement a method that accepts 3 integer values a, b, c. The method should return true if a triangle can be built with the sides of given length and false in any other case.
// (In this case, all triangles must have surface greater than 0 to be accepted).
package main

import "fmt"

func main() {

	fmt.Println(IsTriangle(10, 25, 10)) // false
	fmt.Println(IsTriangle(10, 15, 10)) // true

}

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && a+c > b
}
