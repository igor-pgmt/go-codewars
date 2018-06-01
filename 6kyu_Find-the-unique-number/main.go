// https://www.codewars.com/kata/585d7d5adb20cf33cb000235
// Find the unique number
// There is an array with some numbers. All numbers are equal except for one. Try to find it!
// findUniq([ 1, 1, 1, 2, 1, 1 ]) === 2
// findUniq([ 0, 0, 0.55, 0, 0 ]) === 0.55
// It’s guaranteed that array contains more than 3 numbers.
// The tests contain some very huge arrays, so think about performance.
// This is the first kata in series:
//     Find the unique number (this kata)
//     Find the unique string https://www.codewars.com/kata/585d8c8a28bc7403ea0000c3
//     Find The Unique https://www.codewars.com/kata/5862e0db4f7ab47bed0000e5

package main

import "fmt"

func main() {

	fmt.Println(FindUniq([]float32{1, 1, 1, 2, 1, 1})) // 2
	fmt.Println(FindUniq([]float32{0, 0, 0.55, 0, 0})) // 0.55

}

func FindUniq(arr []float32) float32 {
	if arr[0] != arr[1] && arr[1] == arr[2] {
		return arr[0]
	}
	for key, val := range arr[1:] {
		if val != arr[key] {
			return val
		}
	}
	return 0
}
