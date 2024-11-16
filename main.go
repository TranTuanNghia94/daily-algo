package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{3, 2, 3, 2, 3, 2}
	k := 2
	r := algo.ResultsArray(s, k)
	fmt.Println("\n", r)
}
