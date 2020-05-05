package main

import "fmt"

/*
	Pattern: number of 1s that appear in the binary form of the decimal
	1 = 1 (01)
	2 = 1 (10)
	3 = 2 (11)

	Can count all 1s by iterating through the number
	Example:
	n := 0100011
	count := 0

`	1: count += n & 1, checks to see if first number is a 1
	2: n >>= 1, right shifts to get rid of the 1

	we do this while n != 0

	Since length of binary representation increases logarithmically as decimal increases,
	time complexity is O(logn)
*/

func countIter(n int) int {
	count := 0

	for n != 0 {
		count += n & 1
		n >>= 1
	}

	return count
}

func main() {
	fmt.Println(countIter(102)) // Prints 4
	fmt.Println(countIter(95)) // Prints 6
	fmt.Println(countIter(72)) // Prints 2
	fmt.Println(countIter(60)) // Prints 4
}
