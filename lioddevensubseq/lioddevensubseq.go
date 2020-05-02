package main

import "fmt"

func lioes(arr []int) int {
	n := len(arr)
	var lioes [1000]int
	maxLen := 0

	for i := 0; i < n; i++ {
		lioes[i] = 1
	}

	// Calculate all lioess starting from each index
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] && (arr[i]+arr[j])%2 != 0 && lioes[i] < lioes[j]+1 {
				lioes[i] = lioes[j] + 1
			}
		}
	}

	for i := 0; i < n; i++ {
		if maxLen < lioes[i] {
			maxLen = lioes[i] // Set maxLen to longest subsequence
		}
	}

	return maxLen
}

func main() {
	arr := []int{1, 2, 3, 6, 7, 3}
	fmt.Println(lioes(arr)) // Returns 5 (1, 2, 3, 6, 7)

	arr = []int{4, 3, 8, 10, 0}
	fmt.Println(lioes(arr)) // Returns 2 (3, 8)

	arr = []int{5, 4, 3, 2, 1, 0}
	fmt.Println(lioes(arr)) // Returns 1 (0)
}
