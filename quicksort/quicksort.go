package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
	At first, I initialised an array of size 23 (because I read the text
	file and counted the number of lines) using var arr [23]int. Then I
	got an error message saying that you cant pass an array of type [23]int
	into a function that takes a parameter of type []int. What the hell?
	
	I then changed the parameters of the functions to take in params of type
	[23]int instead, but after 5 minutes I realised that that was a stupid way
	to solve the problem.
	
	Finally I realised that i could initialise an array of type []int and append
	values to it using a function I wrote to extend the length of the array by
	making a copy of it. I wonder if there's a more memory efficient way to do
	this. Maybe I could grow the array by a larger number each time? Something
	to work on in a later question maybe.
*/

func append(original []int, value int) []int { // Appends values to an array
	target := make([]int, len(original)+1)

	copy(target, original[:len(original)])
	target[len(original)] = value

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

	fmt.Println(arr)
	// Returns [82 28 99 10 54 64 40 50 89 22 5 12 7 17 73 49 96 58 80 81 21 37 98]
	
	fmt.Println(quickSort(arr))
	// Returns [5 7 10 12 17 21 22 28 37 40 49 50 54 58 64 73 80 81 82 89 96 98 99]
}
