package main

import "fmt"

func main() {
	fmt.Println(reverseStringUsingRange("World!"))
}

func reverseStringUsingRange(s string) string {
	var out string

	for i := range s {
		out += string(s[len(s)-1-i])
	}
	return out
}
