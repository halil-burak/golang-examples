package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
The primary mechanism for managing state in Go is communication over channels.
We saw this for example with worker pools. There are a few other options for managing state though.
Here we’ll look at using the sync/atomic package for atomic counters accessed by multiple goroutines.
*/
func main() {
	// We’ll use an unsigned integer to represent our (always-positive) counter.
	var ops uint64

	// A WaitGroup will help us to wait for all the goroutines to finish their work.
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// To atomically increment the counter we use AddUint64, giving it the memory address of our ops counter with the & syntax.
		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}

	// We wait until all the goroutines are done.
	wg.Wait()

	// It's safe to access the ops now since all the goroutines are finished. We know that no goroutine is writing to it.
	fmt.Println("Count:", ops)
}

/* We expect to get 50000 as the output of this program.
   Had we used the non-atomic ops++ to increment the counter,
   We'd likely get a different number at each run.
   The reason is the goroutines would be interfering with each other.
   Moreover, we'd get data race failures if we'd use the -race flag
*/
