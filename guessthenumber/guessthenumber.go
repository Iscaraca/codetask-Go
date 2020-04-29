package main

import (
	"fmt"
)

func binarySearch(l, u, a int) int {
	for l <= u {
		middle := (l + u) / 2
		fmt.Printf("Is your number %d?\n", middle)
		if a == middle {
			fmt.Println("Yes.")
			return middle
		} else if a > middle {
			fmt.Println("No, higher.")
			l = middle + 1
		} else { // a < middle
			fmt.Println("No, lower.")
			u = middle - 1
		}
	}
	return -1
}

func main() {
	lower := 2
	upper := 100
	answer := 3

	binarySearch(lower, upper, answer)
	// Guesses 51, 26, 13, 7, 4, 2, then 3.
}
