package main

import (
	"fmt"
)

func add(x, y int) (result int) { // Takes in two integers as parameters, returns result
	result = x + y
	return
}

func main() {
	fmt.Println(add(5, 4)) // Prints 9
	fmt.Println(add(-7, 0)) // Prints -7
}
