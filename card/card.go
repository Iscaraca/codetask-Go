package main

import "fmt"

func shuffle(n int, command string) []int {
	var result []int
	for i := 0; i < n; i++ {
		result = append(result, i)
	}

	commandByte := []byte(command)

	for _, move := range commandByte {
		temp := make([]int, n-1)
    
		if move == byte('A') {
			copy(temp, result[1:])
			temp = append(temp, result[0])
			result = temp
		} else if move == byte('B') {
			tempIndexOne := result[1]
			copy(temp, append(result[0:1], result[2:]...))
			temp = append(temp, tempIndexOne)
			result = temp
		}
	}

	return result
}

func main() {
	k := 3
	result := shuffle(6, "ABBABA.")

	fmt.Printf("%d %d %d\n", result[k-1], result[k], result[k+1])
	// Prints 3 1 5
}
