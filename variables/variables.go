package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
	fmt.Printf("%T\n", a)

	var d bool = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := 3
	fmt.Println(f)
}
