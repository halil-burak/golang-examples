package main

import (
	"fmt"
	"strings"
)

// Index returns the index of the target string or if no match is found it returns -1.
func Index(strs []string, str string) int {
	for i, v := range strs {
		if v == str {
			return i
		}
	}
	return -1
}

// Include returns true if the target string exists in the slice, false if not.
func Include(strs []string, str string) bool {
	return Index(strs, str) >= 0
}

// Any returns true if any of the strings in the slice complies with the predicate with the function.
func Any(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns true if all elements in the slice meets with the predicate provided with the function.
func All(strs []string, f func(string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f.
func Filter(strs []string, f func(string) bool) []string {
	fstrs := make([]string, 0)
	for _, v := range strs {
		if f(v) {
			fstrs = append(fstrs, v)
		}
	}
	return fstrs
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println("Fruits:", strs)

	fmt.Println("Index of pear:", Index(strs, "pear"))
	fmt.Println("Does the fruits include banana:", Include(strs, "banana"))
	fmt.Println("Any:", Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println("All:", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
}
