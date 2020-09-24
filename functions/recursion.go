package main

import "fmt"

func factor(n int) int {
	if n == 0 {
		return 1
	}
	return n * factor(n-1)
}

func main() {
	fmt.Printf("Factor %d is %d\n", 3, factor(3))
}
