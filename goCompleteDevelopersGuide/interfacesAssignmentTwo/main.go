package main

import "fmt"

func main() {
	t := triangle{
		base:   10,
		height: 5,
	}
	s := square{
		sideLength: 10,
	}
	printArea(s)
	printArea(t)
}

type triangle struct {
	base   float64
	height float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
