package main

import (
	"fmt"
	"math"
)

func Sqrt1(x float64) float64 {	// part one. Calculating 10 times. 
	z := 1.0
	for i := 1; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Println(z)	// print each iteration
	}
	return z
}

func Sqrt(x float64) float64 {	// part two. Change the loop condition to stop once the value has stopped changing

	z := x	// 1.0 or x/2
	var d = z - (z*z-x)/(2*z) // d is adjusted z. 
	for math.Abs(d-z) > 1e-5 {	// checks if the difference is greater than 10^5 which, with z = x, was satisfying 
		z, d = z-(z*z-x)/(2*z), z // update z and d
	}
	return z
}

func main() {
	fmt.Println(Sqrt1(8))
	fmt.Printf("My guess is %v, and the expected is %v", Sqrt2(8), math.Sqrt2(8))
}

