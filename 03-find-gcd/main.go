// Finding Euclid's greatest common denominator.
// Let's say you have two numbers 20 and 8, now to find the GCD of this two, we need to
// find the greatest common number that can devide both of this number so that
// they don't have a remainder
// Mathematically, we solve this problem by deviding the largest number by the smaller
// number and if there is a remainder of the division then the divisor become dividend and the remainder
// becomes divisor and will divide again. This process continues until we find the remainder to be zero.
// When remainder becomes zero, the divisor becomes the GCD of the two numbers.

package main

func main() {
	println("The GCD of 20 and 8 is: ", find_gcd(20, 8))
	println("The GCD of 96 and 36 is: ", find_gcd(96, 36))

}

func find_gcd(number, divisor int) int {
	for divisor != 0 {
		t := number
		number = divisor
		divisor = t % number
	}
	return number
}
