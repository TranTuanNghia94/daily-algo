package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 5, 4, 2, 9, 9, 9}
	k := 3
	r := algo.MaximumSubarraySum(s, k)
	fmt.Println("\n", r)
}
