package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{11, 6}
	n := 6
	// k := 10
	r := algo.MinimizedMaximum(n, s)
	fmt.Println("\n", r)
}
