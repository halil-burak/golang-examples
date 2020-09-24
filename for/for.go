package main

import "fmt"

func main() {
	i := 1

	// The most basic type, with one condition
	// Golang has only the for loop, and here it also works like while loops in other languages
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	// A classic initial/condition/after for loop.
	for i := 7; i < 9; i++ {
		fmt.Println(i)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("will stop now")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := 5; n < 9; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
