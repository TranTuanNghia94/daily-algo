package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	// s := []int{1, 5, 4, 2, 9, 9, 9}
	// k := 3
	s := [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}
	r := algo.MaxEqualRowsAfterFlips(s)
	fmt.Println("\n", r)
}
