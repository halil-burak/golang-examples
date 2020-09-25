package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done!")
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)

	isDone := <-done
	if isDone {
		fmt.Println("DONE.")
	} else {
		fmt.Println("FAILED.")
	}

	/*
		done := make(chan bool, 1)
		go worker(done)
		<-done
	*/
}
