package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println("Sum is ", sum)

	for i, num := range nums {
		fmt.Println("index: ", i, "value: ", num)
	}

	// init a map
	kvmap := map[string]string{"a": "apple", "b": "banana"}

	// range on the key, value pairs of the map
	for k, v := range kvmap {
		fmt.Println("key:", k, "value:", v)
	}

	// range on only the keys of the map
	for k := range kvmap {
		fmt.Println("key:", k)
	}

	// range on only the values of the map
	for _, v := range kvmap {
		fmt.Println("value:", v)
	}
}
