// https://www.codewars.com/kata/586c1cf4b98de0399300001d
// Grasshopper - Terminal game combat function
// Create a combat function that takes the player's current health and the amount of damage recieved, and returns the player's new health. Health can't be less than 0.
package main

import "fmt"

func main() {

	fmt.Println(combat(100, 120))
	fmt.Println(combat(120, 100))

}

func combat(health, damage float64) float64 {
	if health > damage {
		return health - damage
	}
	return 0
}
