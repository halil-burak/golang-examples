package main

import (
	"fmt"
	"time"
)

func f(msg chan string) chan string {
	msg <- "hellow"
	return msg
}

func main() {
	messages := make(chan string)
	go f(messages)
	time.Sleep(time.Second)
	gotMsg := <-messages
	fmt.Println(gotMsg)
}
