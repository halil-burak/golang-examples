package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("appended", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied", s)

	m := s[2:5]
	fmt.Println("first slice op:", m)

	m = s[:5]
	fmt.Println("second slice op:", m)

	m = s[2:]
	fmt.Println("third slice op:", m)

	t := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice declared and initialized in one line:", t)
}
