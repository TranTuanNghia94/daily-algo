package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	// s := []int{1, 5, 4, 2, 9, 9, 9}
	// k := 3
	s := [][]int{{3, 2, 4}, {1, 5, 0}}
	r := algo.SlidingPuzzle(s)
	fmt.Println("\n", r)
}
