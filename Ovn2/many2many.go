package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use different random numbers each time this program is executed.
	rand.Seed(time.Now().Unix())

	const strings = 32
	const producers = 4
	const consumers = 2
	consumedStrings := 0

	before := time.Now()
	ch := make(chan string)
	wgp := new(sync.WaitGroup) // Producer WaitGroup
	wgc := new(sync.WaitGroup) // Consumer WaitGroup

	wgp.Add(producers)
	for i := 0; i < producers; i++ {
		go Produce("p"+strconv.Itoa(i), strings/producers, ch, wgp)
	}
	wgc.Add(consumers)
	for i := 0; i < consumers; i++ {
		go Consume("c"+strconv.Itoa(i), ch, wgc, &consumedStrings)
	}
	wgp.Wait() // Wait for all producers to finish.
	close(ch)  // Close channel as the producers are done sending information.
	wgc.Wait() // Wait for all consumers to finish.
	fmt.Println("time:", time.Now().Sub(before), "amount:", consumedStrings)
}

// Produce sends n different strings on the channel and notifies wg when done.
func Produce(id string, n int, ch chan<- string, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		RandomSleep(100) // Simulate time to produce data.
		ch <- id + ":" + strconv.Itoa(i)
	}
	wg.Done()
}

// Consume prints strings received from the channel until the channel is closed.
func Consume(id string, ch <-chan string, wg *sync.WaitGroup, consumedStrings *int) {
	for s := range ch {
		*consumedStrings++
		fmt.Println(id, "received", s)
		RandomSleep(100) // Simulate time to consume data.
	}
	wg.Done() // Consumer is done when all s on ch is read.
}

// RandomSleep waits for x ms, where x is a random number, 0 ≤ x < n,
// and then returns.
func RandomSleep(n int) {
	time.Sleep(time.Duration(rand.Intn(n)) * time.Millisecond)
}
