package main

import "fmt"

func matrixMultiply(a [][]int, b []int) []int {
	var result = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := 0; i < 10; i++ {
		for k := 0; k < 10; k++ {
			result[i] += a[i][k] * b[k]
		}
	}

	return result
}

var trans = [][]int{
	{0, 0, 0, 0, 1, 0, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 1, 0, 1, 0},
	{0, 0, 0, 0, 0, 0, 0, 1, 0, 1},
	{0, 0, 0, 0, 1, 0, 0, 0, 1, 0},
	{1, 0, 0, 1, 0, 0, 0, 0, 0, 1},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{1, 1, 0, 0, 0, 0, 0, 1, 0, 0},
	{0, 0, 1, 0, 0, 0, 1, 0, 0, 0},
	{0, 1, 0, 1, 0, 0, 0, 0, 0, 0},
	{0, 0, 1, 0, 1, 0, 0, 0, 0, 0}}

func calcPaths(start, moves int) int {
	var cumulative = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := 0; i < moves; i++ {
		cumulative = matrixMultiply(trans, cumulative)
	}

	return cumulative[start]
}

func main() {
	fmt.Println(calcPaths(6, 2)) // Returns 6
	fmt.Println(calcPaths(5, 5)) // Returns 0 as there is nowhere to go from 5
	fmt.Println(calcPaths(0, 5)) // Returns 64
}
