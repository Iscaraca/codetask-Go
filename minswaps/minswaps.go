package main

import (
	"fmt"
	"sort"
)

type pair struct {
	value, index interface{}
}

func minSwaps(arr []int) int {
	visited := make([]int, len(arr))
	sorted := make([]pair, len(arr))

	for i, value := range arr {
		sorted[i] = pair{value, i}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].value.(int) <= sorted[j].value.(int)
	})

	count := 0

	for i := 0; i < len(sorted); i++ {
		j := i
		for visited[j] != 1 {
			visited[j] = 1
			j = sorted[j].index.(int)
			count++
			if visited[j] == 1 {
				count--
			}
		}
	}

	return count
}

func main() {
	arr := []int{2, 4, 5, 1, 3}
	fmt.Println(minSwaps(arr)) // Prints 3

	arr = []int{4, 3, 2, 1}
	fmt.Println(minSwaps(arr)) // Prints 2
}
