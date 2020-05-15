package main

import "fmt"

func binaryCoin(n int) int {
  c := 0
  
  for n != 0 {
      n -= (n & (-n))
      c++
  }
  
  return c
}

func main() {
  fmt.Println(binaryCoin(19)) // Prints 3
}
