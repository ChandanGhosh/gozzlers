// Finding max number from an array in an recursive way. Complexity O(n)
package main

import "fmt"

func main() {
	items := []int{2, 4, 2, 6, 76, 32, 98, 67, 87, 21}

	fmt.Printf("Items: %+v\n", items)
	println("max number in the items: ", find_max(items))
}

func find_max(items []int) int {
	if len(items) == 1 {
		return items[0]
	}

	item1 := items[0]
	item2 := find_max(items[1:])

	if item1 > item2 {
		return item1
	} else {
		return item2
	}
}
