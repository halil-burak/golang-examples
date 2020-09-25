package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

// todo add a function that gets a pointer to the channel and move <- and -> ops to this function
// or the function accepts an argument of type channel
