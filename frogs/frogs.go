package main

import (
	"fmt"
	"sort"
)

type pair struct {
	index, number interface{}
}

func frog(xAxis []int, a, b, k int) bool {
	a--
	b--

	sorted := make([]pair, len(xAxis))
	for i, value := range xAxis {
		sorted[i] = pair{i, value}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].number.(int) <= sorted[j].number.(int)
	})

	inRange := false

	for i := range sorted {
		if inRange == false {
			if sorted[i].index.(int) == a || sorted[i].index.(int) == b {
				inRange = true
			}
		} else {
			if i == 0 {
				continue
			} else if sorted[i].number.(int)-sorted[i-1].number.(int) > k {
				return false
			} else if sorted[i].index.(int) == a || sorted[i].index.(int) == b {
				return true
			}
		}
	}

	return false
}

func main() {
	k := 3
	frogs := []int{0, 3, 8, 5, 12}
	fmt.Println(frog(frogs, 1, 2, k)) // Prints true
	fmt.Println(frog(frogs, 1, 3, k)) // Prints true
	fmt.Println(frog(frogs, 2, 5, k)) // Prints false
}
