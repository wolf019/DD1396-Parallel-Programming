package main

import "fmt"

func fibonacci() func() int {
	fibo, fibo0, fibo1 := 0, 0, 1 // return value, first, and second fib.nr
	return func() int {
		fibo = fibo0                      // fibo is updated with the desired fib.nr ro return
		fibo0, fibo1 = fibo1, fibo0+fibo1 // update fib.nr pair.
		return fibo						  // return desired fib.nr
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
