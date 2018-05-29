// https://www.codewars.com/kata/n-point-crossover/go

package main

import "fmt"

func main() {
	a, b := Crossover1([]int{1, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(a, b)
	c, d := Crossover2([]int{1, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	fmt.Println(c, d)

}

// This is the fastest function
func Crossover1(ns []int, xs []int, ys []int) ([]int, []int) {
	resMap := make(map[int]bool)
	for _, v := range ns {
		resMap[v] = true
	}

	flag := false
	for i := 0; i < len(xs); i++ {
		if resMap[i] {
			flag = !flag
		}
		if flag {
			xs[i], ys[i] = ys[i], xs[i]
		}
	}
	return xs, ys
}

func Crossover2(ns []int, xs []int, ys []int) ([]int, []int) {
	x := make(map[int]bool)
	for _, val := range ns {
		x[val] = true
	}
	for k := range x {
		for i := k; i < len(xs); i++ {
			xs[i], ys[i] = ys[i], xs[i]
		}
	}
	return xs, ys
}
