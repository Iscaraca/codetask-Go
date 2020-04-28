package main

import "fmt"

func gcd(a, b int) int { // Where a > b
	if b != 0 {
		return gcd(b, a%b)
	} else {
		return a
	}
}

func main() {
	fmt.Println(gcd(12, 8)) // Prints 4
}
