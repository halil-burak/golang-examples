package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 43
	return &p
}

func main() {
	p := newPerson("ali kaya")
	fmt.Println("Person 1:", p)

	p2 := person{"john", 20}
	fmt.Println(p2)

	p3 := person{name: "will", age: 44}
	fmt.Println(p3)

	s := person{name: "sean conor", age: 66}
	fmt.Println(s)

	sp := &s
	fmt.Println(sp.age)

	// structs are mutable
	sp.age = 88
	fmt.Println(s)
	fmt.Println(s, *sp)
}
