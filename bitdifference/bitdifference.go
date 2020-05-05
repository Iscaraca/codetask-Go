package main

import "fmt"

// Brian Kernighan
// n = n & (n - 1) removes the rightmost set bit of n
func flipBits(a, b int) int {
	flip := a ^ b
	count := 0

	for flip != 0 {
		flip = flip & (flip - 1)
		count++
	}

	return count
}

func main() {
	fmt.Println(flipBits(10, 20)) // Prints 4
}
