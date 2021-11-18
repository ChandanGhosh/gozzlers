// In this problem we shall find the LCD(Lowest/least common denominator) of the two numbers.
// Generally we find LCD in case of fractional number for computation. The LCD helps
// finding the average, adding, subtracting etc. by removing fractions.
// Mathematically, when we have two fractional numbers let's say 1/5 and 5/6, we find the prime factors of both and then
// multiply them removing the duplicates. For example, here we shall find the LCD of 5 and 6. Now prime multiple of 5 is 5*1 or 5
// prime multiple of 6 is 2*3. So to find the LCD we will multiply 2*3*5. So the LCD of both would be 30.
// We can find the LCD by using the formula as LCD = (a*b)/GCD(a,b)

package main

import (
	"flag"
	"fmt"
	"os"
)

func findLCD(a, b int) int {
	return (a * b) / findGCD(a, b)
}

// findGCD shows another way of finding GCD recursively
func findGCD(a, b int) int {
	if a == 0 {
		return b
	}
	return findGCD(b%a, a)
}

func main() {

	if len(os.Args) > 1 {
		a, b := flag.Int("a", 10, "enter a"), flag.Int("b", 12, "enter b")
		flag.Parse()
		fmt.Printf("LCD of %v and %v is %v\n", *a, *b, findLCD(*a, *b))
	} else {
		println("Enter two numbers to find LCD")
	}

}
