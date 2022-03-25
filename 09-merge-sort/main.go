// Here, in Merge sort we use the divide and conquer method. Usually the complexity is O(nlogn) or
// linear logarithmic time. So given you have some
// array, could be unsorted, then in Merge sort we divide the array until each array contains
// single element which then we can consider as sorted as the arrays have only one element each
// then taking two arrays as a and b and merge back in to aother array and continue until all the arrays are taken care(conquer).
// In this process there could be some arrays with more elements than the others while merging back, in that case those array should
// be appended at last

package main

import "fmt"

func main() {
	unsorted := []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9, -9, 10, 7, 0, -1}
	fmt.Printf("Unsorted: %+v\n", unsorted)
	// fmt.Printf("Sorted: %+v", merge_sort(unsorted))
	fmt.Printf("%+v\n", divide(unsorted))
}

func divide(dataset []int) []int {

	if len(dataset) < 2 {
		return dataset
	}

	mid := len(dataset) / 2
	left := divide(dataset[:mid])
	right := divide(dataset[mid:])

	return conquer(left, right)

}

func conquer(left, right []int) []int {
	sorted := []int{}
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i++
		} else {
			sorted = append(sorted, right[j])
			j++
		}
	}

	// only left array values
	for i < len(left) {
		sorted = append(sorted, left[i])
		i++
	}

	// Only right
	for j < len(right) {
		sorted = append(sorted, right[j])
		j++
	}

	return sorted
}
