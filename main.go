package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{0, 1, 2, 2, 5, 7}
	k := 2
	r := algo.GetMaximumXor(s, k)
	fmt.Println("\n", r)
}
