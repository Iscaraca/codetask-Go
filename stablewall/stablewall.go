package main

import "fmt"

func cd(n, k int, arr []int) int {
	currCount := 0
	kCount := 0

	for j := n - 1; j >= 0; j-- {
		if arr[j] == currCount+1 {
			currCount++
			if currCount == k {
				kCount++
				currCount = 0
			}
		} else if arr[j] == 1 {
			currCount = 1
			if currCount == k {
				kCount++
				currCount = 0
			}
		} else {
			currCount = 0
		}
	}

	return kCount

}

func main() {
	fmt.Println(cd(12, 3, []int{1, 2, 3, 7, 9, 3, 2, 1, 8, 3, 2, 1})) // Prints 2
	fmt.Println(cd(4, 2, []int{101, 100, 99, 98}))                    // Prints 0
	fmt.Println(cd(9, 6, []int{100, 7, 6, 5, 4, 3, 2, 1, 100}))       // Prints 1
}
