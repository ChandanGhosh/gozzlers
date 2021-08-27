// In this puzzle, we are going to get to arrays of characters
// arr1 = ['a', 'b', 'c', 'd'] & arr2 = ['x', 'y', 'z', 'c']
// we need to find the common character, in this case c

package main

import "strconv"

func main() {
	arr1 := []rune{'a', 'b', 'c', 'd'}
	arr2 := []rune{'x', 'y', 'c'}

	findCommonChar(arr1, arr2)
	findCommonCharEfficient(arr1, arr2)
}

// This is the first naive aprach towards the problem in hand
// we are iterating through each array and finding the common one.
// Time complexity O(a*b)
// Please note, here we specified complexity as a*b instead of n^2 because
// the array sizes are different..
// This is not so efficient function.
// Let's see if we can make it better in findCommonCharEfficient function
func findCommonChar(arr1, arr2 []rune) bool {
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if arr1[i] == arr2[j] {
				println("Found the common character", strconv.QuoteRune(arr1[i]))
				return true
			}
		}
	}
	return false
}

// Here we are trying a different approach. Instead of nesting the loops,
// we are first creating a hashmap from the existing array. The reason is, hashmap is extremely fast
// read, insert and delete O(1)
// And then another loop with the other array to just find the common item which would be pretty fast and easy.
// Here we bring down the complexity from O(a*b) to O(a+b) which is more efficient.
func findCommonCharEfficient(arr1, arr2 []rune) bool {
	charMap := make(map[rune]bool)
	for i := 0; i < len(arr1); i++ {
		charMap[arr1[i]] = true
	}

	for _, v := range arr2 {
		if _, ok := charMap[v]; ok {
			println("Found the common character", strconv.QuoteRune(v))
			return true
		}
	}
	return false
}
