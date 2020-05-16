package main

import "fmt"

func check(mid int, nums []int, n, K int) bool {
	count := 0
	sum := 0

	for i := 0; i < n; i++ {
		if nums[i] > mid {
			return false
		}

		sum += nums[i]

		if sum > mid {
			count++
			sum = nums[i]
		}
	}

	count++

	if count <= K {
		return true
	}

	return false
}

func solve(nums []int, n, K int) int {
	start := 1
	end := 0

	for i := 0; i < n; i++ {
		end += nums[i]
	}

	answer := 0
	for start <= end {
		mid := (start + end) / 2

		if check(mid, nums, n, K) {
			answer = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return answer
}

func marbles(m int, nums []int) int {
	n := len(nums)
	return solve(nums, n, m)
}

func main() {
	nums := []int{4, 2, 4, 5, 1, 1}
	fmt.Println(marbles(3, nums)) // Prints 7
}
