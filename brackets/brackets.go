package main

import "fmt"

func valid(brackets string) string {
	var stack []byte
	brackByte := []byte(brackets)

	for i := 0; i < len(brackByte); i++ {
		c := brackByte[i]

		if (c == 40) || (c == 91) || (c == 123) {
			stack = append(stack, c)
		} else {
			if len(brackByte) == 0 {
				return "Invalid"
			} else if (stack[len(stack)-1]+1 == c) || (stack[len(stack)-1]+2 == c) {
				stack = stack[:len(stack)-1]
			} else {
				return "Invalid"
			}
		}
	}

	if len(stack) == 0 {
		return "Valid"
	}

	return "Invalid"
}

func main() {
	fmt.Println(valid("()[]{}")) // Prints Valid
	fmt.Println(valid("{[]()}")) // Prints Valid
	fmt.Println(valid("{([)]}")) // Prints Invalid
}
