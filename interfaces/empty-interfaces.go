package main

import "fmt"

// MyString custom string
type MyString string

// Car as vehicle
type Car struct {
	make     string
	maxSpeed int
}

func explain(i interface{}) {
	fmt.Printf("Type of the passed argument is '%T' with the value of %v\n", i, i)
}

// GetCarName getter method for the make of the car
func (c Car) GetCarName() string {
	return c.make
}

// SetCarName setter method for the name of the car
func (c *Car) SetCarName(s string) {
	c.make = s
}

func main() {
	var s MyString = "hey"
	explain(s)
	car := Car{"ford mustang", 320}
	explain(car)
	fmt.Println(car.GetCarName())
	car.SetCarName("mercedes glb")
	explain(car)
	fmt.Println(car.GetCarName())
}
