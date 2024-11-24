package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	// s := []int{1, 5, 4, 2, 9, 9, 9}
	// k := 3
	s := [][]int{{1, -1}, {-1, 1}}
	r := algo.MaxMatrixSum(s)
	fmt.Println("\n", r)
}
