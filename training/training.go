package main

import (
	"fmt"
	"sort"
)

func append(arr []int, value int) []int {
	target := make([]int, len(arr)+1)

	copy(target, arr[:len(arr)])
	target[len(arr)] = value

	return target
}

// length = 4
// Number = 3
// temp = [3,4,5,6]
// dp = [    3,3]

func training(arr []int, number int) int {
	temp := make([]int, len(arr))
	copy(temp, arr)

	sort.Slice(temp, func(i, j int) bool { return temp[i] <= temp[j] })

	dp := make([]int, len(arr)-(number-1)) // Tracks subsequence of length number with the least difference from the last number in subseq

	for i := 0; i < len(dp); i++ {
		difference := temp[i+number-1] * (number - 1)
		for j := i; j < i+number-1; j++ {
			difference -= temp[j]
		}
		dp[i] = difference
	}

	hours := 1000000000

	for _, diff := range dp {
		if diff < hours {
			hours = diff
		}
	}

	return hours
}

func main() {
	var students []int
	input1 := []int{3, 1, 9, 100}
	input2 := []int{5, 5, 1, 2, 3, 4}
	input3 := []int{7, 7, 1, 7, 7}

	for _, skill := range input1 {
		students = append(students, skill)
	}

	fmt.Println(training(students, 3)) // Returns 14

	students = nil
	for _, skill := range input2 {
		students = append(students, skill)
	}

	fmt.Println(training(students, 2)) // Returns 0

	students = nil
	for _, skill := range input3 {
		students = append(students, skill)
	}

	fmt.Println(training(students, 5)) // Returns 6
}
