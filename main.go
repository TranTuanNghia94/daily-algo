package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{2, 2, 2, 1, 1, 1}
	// k := 10
	r := algo.FindLengthOfShortestSubarray(s)
	fmt.Println("\n", r)
}
