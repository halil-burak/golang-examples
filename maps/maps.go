package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println("empty map:", m)

	m["k1"] = 7
	m["k2"] = 12
	fmt.Println("map set:", m)
	fmt.Println("length of the map:", len(m))

	delete(m, "k2")
	fmt.Println("map after delete:", m)

	v1 := m["k1"]
	fmt.Println("Got the value of key k1:", v1)

	_, isK2Present := m["k2"]
	fmt.Println("the key k2 exists?:", isK2Present)

	t := map[string]int{"k3": 4, "k4": 9}
	fmt.Println("map declared and initialized in one line:", t)
}
