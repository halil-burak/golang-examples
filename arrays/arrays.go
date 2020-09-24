package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)
	fmt.Println(len(a))

	b := [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("two d array:", twoD)
}
