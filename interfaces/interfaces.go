package main

import (
	"fmt"
	"math"
)

// Shape interface for calculations about shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangular struct type to implement the Shape interface
type Rectangular struct {
	width  float64
	heigth float64
}

// Circle struct type to implement the Shape interface
type Circle struct {
	radius float64
}

// Area method to implement the Shape interface's Area method (@Override in Java)
func (r Rectangular) Area() float64 {
	return r.width * r.heigth
}

// Perimeter method to implement the Shape interface's Perimeter method (@Override in Java)
func (r Rectangular) Perimeter() float64 {
	return 2 * (r.width + r.heigth)
}

// Area method to implement the Shape interface's Area method (@Override in Java)
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter method to implement the Shape interface's Perimeter method (@Override in Java)
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.radius
}

func main() {
	var s Shape
	fmt.Println("Value is", s)
	fmt.Printf("Type is %T\n", s)
	s = Rectangular{5.0, 4.0}
	fmt.Println("Value is", s, "after value assignment.")
	fmt.Printf("Type is %T\n", s)
	fmt.Println("Type assertion:", s.(Rectangular))

	r := Rectangular{5.0, 4.0}
	fmt.Println("Area of the rectangular is", r.Area())
	fmt.Println("Perimeter of the rectangular is", r.Perimeter())

	c := Circle{10}
	fmt.Println("Area of the circle is", c.Area())
	fmt.Println("Perimeter of the circle is", c.Perimeter())
	s = c
	fmt.Println("Value is", s, "after Circle assignment.")
	fmt.Printf("Type is %T\n", s)
}
