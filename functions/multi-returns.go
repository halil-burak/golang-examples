package main

import "fmt"

func vals() (int, int) {
	return 3, 5
}

func addProduct(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := vals()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	_, c := vals()
	fmt.Println("c:", c)

	m, n := addProduct(2, 7)
	fmt.Println("sum:", m, "product:", n)
}
