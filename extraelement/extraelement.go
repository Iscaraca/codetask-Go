package main

import (
	"fmt"
)

func findDiff(a, b []int, n int) int {
	lower := 0
	upper := n - 1
	index := -1

	for lower <= upper {
		mid := (lower + upper) / 2
		if a[mid] == b[mid] {
			lower = mid + 1
		} else {
			index = mid
			upper = mid - 1
		}
	}

	return index
}

func main() {
	fmt.Println(findDiff([]int{2, 4, 6, 8, 9, 10, 12}, []int{2, 4, 6, 8, 10, 12}, 7)) //Prints 4
	fmt.Println(findDiff([]int{3, 5, 7, 9, 11, 13}, []int{3, 5, 7, 11, 13}, 6)) // Prints 3
}
