package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend.", time.Now().Weekday())
	default:
		fmt.Println("It's a weekday.", time.Now().Weekday())
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.", t.Hour())
	default:
		fmt.Println("It's afternoon.", t.Hour())
	}
	// Type switch
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it's a boolean")
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(1.0)
	whatAmI(1e2)

	var x interface{} = "foo"
	s, ok := x.(string)
	fmt.Println(s, ok)
	n, nok := x.(int)
	fmt.Println(n, nok)
	fmt.Printf("Type of interface is %T\n", x)

	var m interface{} = 23
	k, isConverted := m.(string)
	fmt.Println(k, isConverted)
	fmt.Printf("Type of interface is %T\n", m)
}
