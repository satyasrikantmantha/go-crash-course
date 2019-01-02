package main

import (
	"fmt"
	"math"
)

//Define interface
type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 10}
	rectangle := Rectangle{10, 5}

	fmt.Printf("Area of circle: %f\n", getArea(circle))
	fmt.Printf("Area of rectangle: %f\n", getArea(rectangle))

}
