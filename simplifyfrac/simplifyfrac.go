package main

import "fmt"

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func simplify(a, b int) {
	commonDivisor := gcd(a, b)
	fmt.Printf("%d/%d\n", a/commonDivisor, b/commonDivisor)
}

func main() {
	simplify(2, 4) // Prints 1/2
	simplify(332, 874) // Prints 166/437
	simplify(3, 3) // Prints 1/1
}
