/*
Given a sorted array and a number x, find a pair in an array whose sum is closest to x.

Examples:

Input: arr[] = {10, 22, 28, 29, 30, 40}, x = 54
Output: 22 and 30

Input: arr[] = {1, 3, 4, 7, 10}, x = 15
Output: 4 and 10

Complexity : O(n)

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	goal := 54
	// goal := 15
	arr := []int{10, 22, 28, 29, 30, 40}
	// arr := []int{1, 3, 4, 7, 10}
	arrlen := len(arr)

	fmt.Printf("Given sorted array: %+v and Goal is %v\n", arr, goal)
	findClosest(arr, goal, arrlen)
}

func findClosest(arr []int, goal int, arrlen int) {
	leftIndx := 0
	rightIndx := arrlen - 1
	diff := math.MaxInt32
	var diffLocal, num1, num2 int

	for rightIndx > 1 && rightIndx != leftIndx {
		total := arr[leftIndx] + arr[rightIndx]
		if total <= goal {
			diffLocal = goal - total
			if diffLocal < diff {
				diff = diffLocal
				num1 = arr[leftIndx]
				num2 = arr[rightIndx]
			}
			leftIndx++
		} else {
			diffLocal = total - goal
			if diffLocal < diff {
				diff = diffLocal
				num1 = arr[leftIndx]
				num2 = arr[rightIndx]
			}
			rightIndx--
		}

	}

	fmt.Printf("The pair closest to goal %v is: %v %v\n", goal, num1, num2)

}
