package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 2, 8}
	k := 10
	r := algo.MinimumSubarrayLength(s, k)
	fmt.Println("\n", r)
}
