package main

// Add adds the numbers in a and sends the result on res.
func Add(a []int, res chan<- int) {
	sum := 0
	for _, num := range a {
		sum += num
	}
	res <- sum
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	n := len(a)
	ch := make(chan int)
	go Add(a[:n/2], ch)
	go Add(a[n/2:], ch)

	subtot1, subtot2 := <-ch, <-ch // get the two values on the channel
	print(subtot1 + subtot2)       // print sum

}
