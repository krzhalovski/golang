package main

import (
	"fmt"
	"math"
)

// Shape : interface for geometrical shapes
type Shape interface {
	area() float64
}

// Circle : Circle struct
type Circle struct {
	x, y, radius float64
}

// Rectangle : Rectanlge struct
type Rectangle struct {
	x, y, width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{0, 0, 2}
	r := Rectangle{0, 0, 2, 2}

	fmt.Println(getArea(c))
	fmt.Println(getArea(r))
}
