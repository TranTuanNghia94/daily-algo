package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 1, 1, 1, 1}
	k := 2
	r := algo.MaximumSubarraySum(s, k)
	fmt.Println("\n", r)
}
