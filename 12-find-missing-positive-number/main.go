/*
Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.

*/

package main

import "sort"

func main() {
	// arr := []int{3, 4, -1, 1}
	arr := []int{-2, 3, 4, -1, 2, 1, 6}
	// arr := []int{1, 2, 0}
	// arr := []int{1, 2, 0, 3, 5}

	println(findMissingPositive(arr))
}

// Complexity - O(n) T, O(1) S
func findMissingPositive(arr []int) int {

	sort.Ints(arr) // O(n) T

	first, last := 0, len(arr) //O(1) S
	for first+1 < last {       // O(n) T
		if arr[first] > 0 && arr[first]+1 != arr[first+1] {
			return arr[first+1] - 1
		}
		first++
	}

	return arr[last-1] + 1
}
