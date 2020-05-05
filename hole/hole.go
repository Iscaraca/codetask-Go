package main

import (
	"fmt"
	"math"
)

func isSafe(i, j int, field [][]int) bool {
	if i < 0 || i >= len(field) || j < 0 || j >= len(field[i]) {
		return false
	}
	return true
}

func dfs(i, j, count int, field [][]int) int {
	//check diagonal right is 1 or not
	if !isSafe(i+1, j+1, field) {
		return count
	}
	//then check for max size of square formed
	for k := 0; k <= count; k++ {
		if field[i+1-k][j+1] == 1 || field[i+1][j+1-k] == 1 {
			return count
		}
	}

	return dfs(i+1, j+1, count+1, field)
}

func maxSquare(field [][]int) int {
	n := len(field)
	if n == 0 {
		return 0
	}

	m := len(field[0])
	ans := 0

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == 0 {
				ans = int(math.Max(float64(ans), float64(dfs(i, j, 1, field))))
			}
		}
	}
	return ans
}

func main() {
	var field [][]int
	field = [][]int{
		{1, 0, 1, 0, 1, 0},
		{0, 1, 0, 1, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 0, 0}}

	fmt.Println(maxSquare(field)) // Prints 3

	field = [][]int{
		{0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0}}

	fmt.Println(maxSquare(field)) // Prints 5
}
