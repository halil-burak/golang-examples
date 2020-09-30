package main

import (
	"fmt"
	"sort"
)

func main() {
	ints := []int{2, 4, 1, 43}
	sort.Ints(ints)
	fmt.Println("Sorted integers:", ints)
	strs := []string{"d", "b", "c", "a"}
	sort.Strings(strs)
	fmt.Println("Sorted strings:", strs)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Integers are sorted:", s)
}
