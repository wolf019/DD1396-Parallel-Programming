package main

import (
	"fmt"
	"sync"
)

// This program should go to 11, but sometimes it only prints 1 to 10.
// See readme for solution.
func main() {

	ch := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go Print(ch, &wg) // We want to wait until Print is done. Which happens when channel is closed and all values on the channel is handled by range.

	for i := 1; i <= 11; i++ {
		ch <- i
	}

	close(ch) // No more values will be sent on ch.
	wg.Wait() // Wait for Print to finish.
}

// Print prints all numbers sent on the channel.
// The function returns when the channel is closed.
func Print(ch <-chan int, wg *sync.WaitGroup) {
	for n := range ch { // reads from channel until it's closed (until all values have been read.)
		fmt.Println(n)
	}
	wg.Done() 
}
