package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func append(original []int, value int) []int { // Appends values to an array
	target := make([]int, len(original)+1)

	copy(target, original[:len(original)])
	target[len(original)] = value
	copy(target[len(original)+1:], original[len(original):])

	return target
}

func quickSort(arr []int) []int {
	newArr := make([]int, len(arr))

	for i, v := range arr {
		newArr[i] = v
	}

	sort(newArr, 0, len(arr)-1)

	return newArr
}

func sort(arr []int, start, end int) {
	if (end - start) < 1 { // When partitions are too small to be sorted
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot { // Moves numbers smaller than pivot to the start
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++ // Keeps track of where the pivot will end up
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	sort(arr, start, splitIndex-1) // Sort the segment behind pivot
	sort(arr, splitIndex+1, end) // Sort the segment after pivot
}

func main() {
	var arr []int
	readFile, err := os.Open("data.txt") // Reads from file
	if err != nil {
		fmt.Println("Error in opening data file.")
	}
	defer readFile.Close()

	scanner := bufio.NewScanner(readFile)

	for scanner.Scan() {
		new, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, new) // Append values by line to arr[]
	}

	readFile.Close()

	fmt.Println(arr) // Returns [82 28 99 10 54 64 40 50 89 22 5 12 7 17 73 49 96 58 80 81 21 37 98]
	fmt.Println(quickSort(arr)) // Returns [5 7 10 12 17 21 22 28 37 40 49 50 54 58 64 73 80 81 82 89 96 98 99]
}
