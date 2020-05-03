// My attempt at 100 Doors using Go
package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ { // For every door in sequence
		count := 0
		for j := 1; j < 101; j++ {
			if i%j == 0 { // I calculate how many times the state of that door changes throughout the 100 passes
				count++
			}
		}
		if count%2 == 1 { // If the number of state changes is odd, that means the door will be open since all doors are initially closed
			fmt.Println("Door", i, "is open.")
		}
	}
}
