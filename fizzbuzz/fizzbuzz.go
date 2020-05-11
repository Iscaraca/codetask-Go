package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		c := 0
		if i%3 == 0 {
			fmt.Printf("fizz")
			c++
		}
		if i%5 == 0 {
			fmt.Printf("buzz")
			c++
		}
		if c == 0 {
			fmt.Printf("%d", i)
		}
		fmt.Printf("\n")
	}
}
