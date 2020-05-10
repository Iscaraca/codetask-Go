package main

import (
	"fmt"
	"math"
)

func memsetButInGo(arr [][]int) {
	for i := range arr {
		for j := range arr[i] {
			arr[i][j] = -1
		}
	}
}

func knapSack(capacity, n int, value, weight []int, dp [][]int) int {
	// Subproblem alr solved
	if dp[capacity][n] != -1 {
		return dp[capacity][n]
	} else if n == 0 || capacity == 0 {
		dp[capacity][n] = 0
		return 0
	} else if weight[n-1] > capacity {
		dp[capacity][n] = knapSack(capacity, n-1, value[:n-1], weight[:n-1], dp)
		return dp[capacity][n]
	} else {
		dp[capacity][n] = int(math.Max(float64(knapSack(capacity, n-1, value[:n-1], weight[:n-1], dp)), float64(value[n-1]+knapSack(capacity-weight[n-1], n-1, value[:n-1], weight[:n-1], dp))))
		return dp[capacity][n]
	}
}

func main() {
	capacity := 20
	n := 4
	dp := make([][]int, capacity+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	memsetButInGo(dp)

	value := []int{10, 40, 50, 75}
	weight := []int{5, 10, 3, 12}

	fmt.Println(knapSack(capacity, n, value, weight, dp)) // Prints 135 (75 + 50 + 10, left with a capacity of 0)
}
