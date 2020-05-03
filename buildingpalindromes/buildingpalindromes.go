package main

import "fmt"

func palindrome(str string, lower, upper int) bool {
	slice := str[lower-1 : upper]
	strByte := []byte(slice)

	var occurrences [256]int

	for i := 0; i < len(strByte); i++ {
		occurrences[strByte[i]]++
	}

	odd := 0

	for _, occurrence := range occurrences {
		if occurrence%2 != 0 {
			odd++
		}

		if odd > 1 { // Palindromes have either 1 odd occuring character in the middle or no odd occurring characters
			return false
		}
	}
	return true
}

func main() {
	str := "ABAACCA"
	fmt.Println(palindrome(str, 3, 6)) // Prints true, AACC -> ACCA
	fmt.Println(palindrome(str, 4, 4)) // Prints true, A is already a palindrome
	fmt.Println(palindrome(str, 2, 5)) // Prints false, BAAC cannot be rearranged to form a palindrome
}
