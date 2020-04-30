package main

import (
	"fmt"
	"math"
)

type Shape struct {
	area      float64
	perimeter float64
}

type Square struct {
	length float64
	Shape  // Composition by embedding struct Shape, sort of like inheritance in a way but hierarchy doesn't exist
}

type equilateralTriangle struct {
	base  float64
	Shape // Composition by embedding struct Shape
}

func (s *Square) setSquDetails() {
	s.area = s.length * s.length
	s.perimeter = s.length * 4 
}

func (t *equilateralTriangle) setTriDetails() {
	t.area = t.base * t.base * (math.Sqrt(3) / 4)
	t.perimeter = t.base * 3
}

func main() {
	GenericShape := Shape{0, 0} // Generic shape struct for later shapes to inherit from

	Square1 := Square{5, GenericShape} // Object Square1 with length 5
	Square1.setSquDetails()
	fmt.Printf("A square with a side length of %.1f has an area of %.1f ", Square1.length, Square1.area)
	fmt.Printf("and a perimeter of %.1f\n", Square1.perimeter)
	// A square with a side length of 5.0 has an area of 25.0 and a perimeter of 20.0

	Triangle1 := equilateralTriangle{5, GenericShape} // Object Triangle1 with base 5
	Triangle1.setTriDetails()
	fmt.Printf("An equilateral triangle with a side length of %.1f has an area of %.1f ", Triangle1.base, Triangle1.area)
	fmt.Printf("and a perimeter of %.1f\n", Triangle1.perimeter)
	// An equilateral triangle with a side length of 5.0 has an area of 10.8 and a perimeter of 15.0
}
