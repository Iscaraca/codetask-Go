package main

import "fmt"

func main() {
	s := 76
	c := 0

	for s >= 1000000 {
		c += s / 1000000
		s = s % 1000000
	}
	for s >= 500000 {
		c += s / 500000
		s = s % 500000
	}
	for s >= 100000 {
		c += s / 100000
		s = s % 100000
	}
	for s >= 50000 {
		c += s / 50000
		s = s % 50000
	}
	for s >= 10000 {
		c += s / 10000
		s = s % 10000
	}
	for s >= 5000 {
		c += s / 5000
		s = s % 5000
	}
	for s >= 1000 {
		c += s / 1000
		s = s % 1000
	}
	for s >= 500 {
		c += s / 500
		s = s % 500
	}
	for s >= 100 {
		c += s / 100
		s = s % 100
	}
	for s >= 50 {
		c += s / 50
		s = s % 50
	}
	for s >= 10 {
		c += s / 10
		s = s % 10
	}
	for s >= 5 {
		c += s / 5
		s = s % 5
	}
	for s >= 1 {
		c += s / 1
		s = s % 1
	}
	fmt.Println(c) // Prints 5
}
