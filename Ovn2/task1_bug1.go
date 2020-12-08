package main

import "fmt"

// I want this program to print "Hello world!", but it doesn't work.
// See readme for solution. 
func main() {
	ch := make(chan string, 1)
	ch <- "Hello world!"
	fmt.Println(<-ch)
}
