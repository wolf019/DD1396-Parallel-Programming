package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	server := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
	}

	// Add a time limit for all requests made by this client.
	client := &http.Client{Timeout: 10 * time.Second}

	for {
		before := time.Now()
		// res := Get(server[1], client)
		res := MultiGet(server, client)
		after := time.Now()
		fmt.Println("Response:", res)
		fmt.Println("Time:", after.Sub(before))
		fmt.Println()
		time.Sleep(500 * time.Millisecond)
	}
}

// Response by Stefan
type Response struct {
	Body       string
	StatusCode int
}

func (r *Response) String() string {
	return fmt.Sprintf("%q (%d)", r.Body, r.StatusCode)
}

// Get makes an HTTP Get request and returns an abbreviated response.
// The response is empty if the request fails.
func Get(url string, client *http.Client) *Response {
	resp, err := client.Get(url)
	if err != nil {
		return &Response{}
	}
	// res.Body != nil when err == nil
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ReadAll: %v", err)
	}
	return &Response{string(body), resp.StatusCode}
}

// AtomicInt helps with data races
type AtomicInt struct {
	mu sync.Mutex // A lock than can be held by one goroutine at a time.
	n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
	a.mu.Lock() // Wait for the lock to be free and then take it.
	a.n += n
	a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
	a.mu.Lock()
	n := a.n
	a.mu.Unlock()
	return n
}

// MultiGet makes an HTTP Get request to each url and returns
// the response from the first server to answer with status code 200.
// If none of the servers answer before timeout, the response is 503
// – Service unavailable.
func MultiGet(urls []string, client *http.Client) *Response {

	respCh := make(chan *Response, 1) // Response channel
	doneCh := make(chan int, 1)       // Done channel

	var noTimeOuts AtomicInt // number of time outs

	for _, url := range urls {
		go func(url string) {
			resp := Get(url, client)
			if resp.StatusCode == 200 { // If we get response we send it to the response channel
				respCh <- resp
			}
			if resp.StatusCode == 503 || resp.StatusCode == 0 { // If server answers with timeout or error.
				noTimeOuts.Add(1)                    // Add one to the number of time outs
				if noTimeOuts.Value() == len(urls) { // if we have time out on all requests, we are done.
					doneCh <- 0
				}
			}
		}(url)
	}

	select {
	case resp := <-respCh:
		return resp // Send response
	case <-doneCh:
		return &Response{} // Send empty response.
	}
}
