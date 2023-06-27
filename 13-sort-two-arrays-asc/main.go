package main

import (
	"fmt"
)

func main() {
	arr1 := []int{100, 23, 12, 54, 999}
	arr2 := []int{34, 21, 7, 5, 31, 67}

	find(arr1, arr2)
}

// arr1 = [100,23,12,54, 999]
// arr2 = [34, 21, 7, 5]
// arr3 = [5, 7, 12, 21, 34, 54, 100, 999] // output

func find(arr1, arr2 []int) {
	var merged []int

	merged = append(merged, arr1...)
	merged = append(merged, arr2...)

	fmt.Printf("%+v ", divideandconquer(merged))

}

func divideandconquer(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	mid := len(arr) / 2
	left := divideandconquer(arr[:mid])
	right := divideandconquer(arr[mid:])

	var i, j int
	var sorted []int

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	if i < len(left) {
		sorted = append(sorted, left[i:]...)
	}

	if j < len(right) {
		sorted = append(sorted, right[j:]...)
	}

	return sorted
}
