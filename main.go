package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := [][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}
	q := []int{1, 2, 3, 4, 5, 6}
	// k := 10
	r := algo.MaximumBeauty(s, q)
	fmt.Println("\n", r)
}
