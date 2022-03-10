// Merge two sorted arrays like [1, 2, 3, 5] and [4, 6, 8, 9]
// Here, in Merge sort we use the divide and conquer method. Usually the complexity is O(nlogn) or
// linear logarithmic time. So given you have some
// array, could be unsorted, then in Merge sort we divide the array until each array contains
// single element which then we can consider as sorted as the arrays have only one element each
// then taking two arrays as a and b and merge back in to aother array and continue until all the arrays are taken care.
// In this process there could be some arrays with more elements than the others while merging back, in that case those array
// elements should be appended as they are already sorted.

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 7}
	b := []int{4, 6, 8, 9}
	fmt.Println(mergeSortedArrays(a, b))
}

//[0, 2, 3, 5] and [4, 6, 8, 9]
func mergeSortedArrays(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	i := 0
	j := 0
	k := 0
	var mergedArr []int
	for i < len(a) && j < len(b) {

		if a[i] < b[j] {
			mergedArr = append(mergedArr, a[i])
			i++

		} else {
			mergedArr = append(mergedArr, b[j])
			j++
		}
		k++
	}

	// Above this point, we had only compared two arrays and sort based on that
	// Now what if left array (a) also have some elements left after the other one, the right array
	// is enumerated or vice versa, we should take care of that scenerios
	// left array elements left
	if i < len(a) {
		mergedArr = append(mergedArr, a[i:]...)
	}

	// right array have elements left
	if j < len(b) {
		mergedArr = append(mergedArr, b[j:]...)
	}

	return mergedArr
}
