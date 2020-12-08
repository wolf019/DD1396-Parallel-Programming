package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {

	countMap := make(map[string]int) // make new map string to int
	words := strings.Fields(s)       // Fields splits s around each instance of one or more consecutive white space characters

	for _, p := range words {
		countMap[p] += 1 // for every string in words add one to the current value
	}

	return countMap
}

func main() {
	wc.Test(WordCount)
}
