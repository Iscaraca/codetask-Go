package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{3, 4, 2, 7, 5}
	sort.Slice(arr, func(i, j int) bool { // Pass anonymous function to sort.Slice for descending comparator
		return arr[i] >= arr[j]
	})

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

}
