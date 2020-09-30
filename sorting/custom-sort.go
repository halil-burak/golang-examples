package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		strs := []string{s[i], s[j]}
		return sort.StringsAreSorted(strs)
	}
	return len(s[i]) < len(s[j])
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fruits := byLength{"peach", "kiwi", "banana", "apple", "pear", "pineapple", "mango", "melon"}
	sort.Sort(fruits)
	fmt.Println("Sorted by length:", fruits)
}
