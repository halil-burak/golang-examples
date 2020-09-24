package main

import "fmt"

// Shape interface
type Shape interface {
	Area() float64
}

// Object interface
type Object interface {
	Volume() float64
}

// Material is combination of two interfaces
type Material interface {
	Shape
	Object
}

// Cube is a shape with 6 identical surfaces
type Cube struct {
	side float64
}

// Area of Shape interface is implemented
func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

// Volume of Object interface is implemented
func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	fmt.Println("Area is", c.Area())
	fmt.Println("Volume is", c.Volume())

	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("Type of m is '%T' and value is '%v'\n", m, m)
	fmt.Printf("Type of m is '%T' and value is '%v'\n", s, s)
	fmt.Printf("Type of m is '%T' and value is '%v'", o, o)
}
