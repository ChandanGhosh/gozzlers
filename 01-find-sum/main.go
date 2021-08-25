// Given a set of numbers let's assume arrays, find if there is pair of numbers which
// sum to 8
// for example: [1, 2, 3, 9] in this array of number there is no pair which can
// sum to 8
// Another example: [1, 2, 4, 4] Here clearly a pair (4,4) there which sum to 8.

package main

import "fmt"

func hasPairWithSum(data []int, sum int) bool {
	// let's create a hashmap(more specifically map in go) to hold the complement of the numbers will traverse.
	// For example when traversing for the first array, the first element is 1 and its complement to sum 8 is (8-1)=7
	// we shall store that in the complement map
	var c map[int]int = make(map[int]int)

	for _, v := range data {
		if num, ok := c[v]; ok {
			fmt.Printf("The pair is %v:%v\n", num, v)
			return true
		} else {
			c[sum-v] = v
		}
	}
	return false
}

func main() {
	a := []int{1, 2, 3, 9}
	hasPairWithSum(a, 8)
	b := []int{1, 2, 4, 4}
	hasPairWithSum(b, 8)
	c := []int{1, 2, 4, 5, 6, 9}
	hasPairWithSum(c, 8)
}
