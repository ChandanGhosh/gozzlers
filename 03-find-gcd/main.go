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
