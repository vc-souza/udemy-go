package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Printf("Area of shape is: %f\n", s.getArea())
}

func main() {
	printArea(square{sideLength: 10.0})
	printArea(triangle{base: 10.0, height: 5})
}
