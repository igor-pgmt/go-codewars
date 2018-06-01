// https://www.codewars.com/kata/51ba717bb08c1cd60f00002f
// Range Extraction
// A format for expressing an ordered list of integers is to use a comma separated list of either
//     individual integers
//     or a range of integers denoted by the starting integer separated from the end integer in the range by a dash, '-'. The range includes all integers in the interval including both endpoints. It is not considered a range unless it spans at least 3 numbers. For example ("12, 13, 15-17")
// Complete the solution so that it takes a list of integers in increasing order and returns a correctly formatted string in the range format.
// Example:
// solution([-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20]);
// returns "-6,-3-1,3-5,7-11,14,15,17-20"

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(Solution([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))

}

func Solution(list []int) string {
	resArr := make([][]string, len(list))
	var result []string
	j := 0
	resArr[0] = append(resArr[0], strconv.Itoa(list[0]))
	for i := 1; i < len(list); i++ {
		if list[i] != list[i-1]+1 {
			j++
		}
		resArr[j] = append(resArr[j], strconv.Itoa(list[i]))
	}

	for _, v := range resArr[:j+1] {
		if len(v) > 2 {
			result = append(result, v[0]+"-"+v[len(v)-1])
		} else {
			result = append(result, strings.Join(v, ","))
		}
	}

	return strings.Join(result, ",")
}
