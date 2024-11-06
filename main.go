package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 2, 2, 3}
	r := algo.IsMonotonic(s)
	fmt.Println("\n", r)
}
