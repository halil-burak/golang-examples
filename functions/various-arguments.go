package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum:", total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	oneslice := []int{4, 5, 6}
	sum(oneslice...)
}
