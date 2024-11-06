package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{3, 16, 8, 4, 2}
	r := algo.CanSortArray(s)
	fmt.Println("\n", r)
}
