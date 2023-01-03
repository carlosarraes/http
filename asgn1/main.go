package main

import "fmt"

type (
	square struct {
		sideLength float64
	}
	triangle struct {
		base, height float64
	}
	shape interface {
		getArea() float64
	}
)

func main() {
	s := square{sideLength: 20}
	t := triangle{base: 20, height: 10}
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return t.base * t.height / 2
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
