package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy) // make double array of unsigned 8-bit integers
	data := make([]uint8, dx) // make array of uint8 

	for i := range pic { // for every "row" in pic
		for j := range data { // for every index in data
			data[j] = uint8(i * j) // assign  
		}
		pic[i] = data // "row" in pic is data
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
