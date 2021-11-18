// Finding Euclid's greatest common denominator.
// Let's say you have two numbers 20 and 8, now to find the GCD of this two, we need to
// find the greatest common number that can devide both of this number so that
// they don't have a remainder
// Mathematically, we solve this problem by deviding the largest number by the smaller
// number and if there is a remainder of the division then the divisor become dividend and the remainder
// becomes divisor and will divide again. This process continues until we find the remainder to be zero.
// When remainder becomes zero, the divisor becomes the GCD of the two numbers.

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		a, b := flag.Int("a", 20, "enter a"), flag.Int("b", 8, "enter b")
		flag.Parse()
		fmt.Printf("The GCD of %v and %v is: %v\n", *a, *b, find_gcd(*a, *b))
	} else {
		println("enter two numbers to find GCD")
	}

}

func find_gcd(number, divisor int) int {
	for divisor != 0 {
		t := number
		number = divisor
		divisor = t % number
	}
	return number
}
