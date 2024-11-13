package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 7, 9, 2, 5}
	lower := 11
	upper := 11
	// k := 10
	r := algo.CountFairPairs(s, lower, upper)
	fmt.Println("\n", r)
}
