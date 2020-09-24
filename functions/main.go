package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	sum := plus(2, 3)
	fmt.Println("2+3:", sum)
	sump := plusPlus(2, 3, 4)
	fmt.Println("2+3+4:", sump)
}
