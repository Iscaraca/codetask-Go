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

func longestPathFromLantern(i, j int, lanterns, dp [][]int) int {
	n := len(lanterns)

	if i < 0 || i >= n || j < 0 || j >= n { // Checks boundaries
		return 0
	}

	if dp[i][j] != -1 { // Subproblem already solved
		return dp[i][j]
	}

	x := -1
	y := -1
	u := -1
	d := -1

	// All values of k are distinct, so there can only be at most one possible way to go from every lantern
	if j < (n-1) && ((lanterns[i][j] - 1) == lanterns[i][j+1]) {
		x = 1 + longestPathFromLantern(i, j+1, lanterns, dp)
	}
	if j > 0 && ((lanterns[i][j] - 1) == lanterns[i][j-1]) {
		y = 1 + longestPathFromLantern(i, j-1, lanterns, dp)
	}
	if i > 0 && ((lanterns[i][j] - 1) == lanterns[i-1][j]) {
		u = 1 + longestPathFromLantern(i-1, j, lanterns, dp)
	}
	if i < (n-1) && ((lanterns[i][j] - 1) == lanterns[i+1][j]) {
		d = 1 + longestPathFromLantern(i+1, j, lanterns, dp)
	}

	// If no lantern is one less than the current one then we take 1 (to represent the length of the path)
	dp[i][j] = int(math.Max(float64(x), float64(math.Max(float64(y), float64(math.Max(float64(u), float64(math.Max(float64(d), float64(1))))))))) // This is so dumb, I wonder if there's a better max function in Go

	return dp[i][j]
}

func longestPath(lanterns [][]int) int {
	result := 1
	n := len(lanterns)

	// Create a lookup table and fill all entries in it as -1. I wonder if there's a better way to make 2D arrays in Go
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	memsetButInGo(dp)

	// Compute longest path beginning from all lanterns
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == -1 {
				longestPathFromLantern(i, j, lanterns, dp)
			}

			// Update result if needed
			result = int(math.Max(float64(result), float64(dp[i][j])))
		}
	}

	return result
}

func main() {
	lanterns := [][]int{{1, 2, 9}, {5, 3, 8}, {4, 6, 7}}
	fmt.Println(longestPath(lanterns)) // Prints 4
}
