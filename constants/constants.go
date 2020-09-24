package main

import (
	"fmt"
	"math"
)

func main() {
	const name = "cats"
	fmt.Println(name)

	const n = 500000
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
