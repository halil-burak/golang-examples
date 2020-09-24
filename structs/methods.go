package main

import "fmt"

type rectangular struct {
	width, heigth int
}

func (r rectangular) area() int {
	return r.width * r.heigth
}

func (r *rectangular) perimeter() int {
	return 2 * (r.width + r.heigth)
}

func main() {
	r := rectangular{3, 4}
	fmt.Println("Area:", r.area(), "Perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
