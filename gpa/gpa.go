package main

import "fmt"

func gpa(a int) float64 {
	c := 0.8

	if a < 40 { // Failing grade
		return 0.8
	} else if a >= 80 { // Full gpa
		return 4.0
	} else if a >= 75 { // 70 - 79 has an interval of 10 while the other grades have an interval of 5, so we can split into two intervals of 5 
		return 3.6
	}

	for a -= 40; a >= 0; a -= 5 { // Uses math to calculate the gpa since unique cases are accounted for and all intervals are now regular
		c += 0.4
	}

	return c
}

func main() {
	fmt.Printf("%.1f\n", gpa(50))  // Returns 2.0
	fmt.Printf("%.1f\n", gpa(49))  // Returns 1.6
	fmt.Printf("%.1f\n", gpa(79))  // Returns 3.6
	fmt.Printf("%.1f\n", gpa(80))  // Returns 4.0
	fmt.Printf("%.1f\n", gpa(100)) // Returns 4.0
}
