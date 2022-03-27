// Quick sort follows the divide and conquer algorithm like merge sort with only one difference. It actually partition the data based
// on a pivot point and continues recursively untill all lesser values in the left and higher values at the right. The pivot position
// can be chosen differently. Some algos choose it as the first index,some as last, some randmom positions etc. Here we are choosing
// the last index. Now, when first index is less than last index, if the array value in start index is lower than the pivot value
// then swap the first and the current index values. That way we ensure, we have values from the start of the arrays which are lesser
// than the pivot. Once we exhaust the array and don't find any more values less than the pivot, we reach at the index where we can now safely
// place the pivot value. At this point, the array to the left of the pivot is lesser values than the pivot and the the values to the right
// of the pivot are higher. Now we can divide the array in to two part in the pivot position and recursively continue partitioning based
// on the above partioning logic. When the recursion completes, we end-up having all the items sorted. Generally this performs better
// as works in spliting the arrays. So the complexity is O(n*log(n)). But in case the array is partially sorted, in that case it takes
// significant time as it stills continue breaking based on partitions and pivot being always lowest or highest point and never in the middle.
// In that case it can go upto Quadratic time complexity - O(n^2). But that is generally not so common case for the elements.

package main

import (
	"fmt"
)

func main() {
	unsorted := []int{20, 6, 8, 53, 23, 87, 42, 19, -1, -99}

	fmt.Printf("%+v\n", quickSort(unsorted, 0, len(unsorted)-1))
}

func quickSort(unsorted []int, low, high int) []int {

	if low < high {
		var partitionIndx int
		unsorted, partitionIndx = partition(unsorted, low, high)
		unsorted = quickSort(unsorted, low, partitionIndx-1)
		unsorted = quickSort(unsorted, partitionIndx+1, high)
	}
	return unsorted
}

func partition(arr []int, first, last int) ([]int, int) {

	pivot := arr[last]
	i := first
	for j := first; j < last; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[last] = arr[last], arr[i]
	return arr, i
}
