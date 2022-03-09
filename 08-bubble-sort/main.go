// Bubble Sort is not the most efficient sorting algorithm
// But it is easier to read and understand so we shall explore
// this to understand it. We shall be working with an array
// [6, 20, 8, 19, 56, 23, 87, 41, 49, 53]
// usually for a selected index, it iterates over all the items and if any item
// is less than the currently selected item, it swaps the values
// and the process continues until all the items are not sorted.
// The time complexity of this algorithm is O(n^2)
package main

import "fmt"

func main() {
	unsorted := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	fmt.Printf("Sorted: %+v\n", bubbleSort(unsorted))
}

func bubbleSort(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	//Quadratic complexity O(n^2)
	for i := 0; i < len(arr); i++ { // O(n)
		for j := i + 1; j < len(arr); j++ { // O(n)
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr

}
