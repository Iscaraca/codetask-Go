package main

import "fmt"

func reversestring(str string) string {
	strbyte := []byte(str)
	i := 0
	j := len(str) - 1
	var temp byte
	for i < j {
		temp = strbyte[j]
		strbyte[j] = strbyte[i]
		strbyte[i] = temp
		i++
		j--
	}

	return string(strbyte)
}

func main() {
	fmt.Println(reversestring("cabbage")) // Prints egabbac
}
