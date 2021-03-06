// Merge two sorted arrays like [1, 2, 3, 5] and [4, 6, 8, 9]

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
