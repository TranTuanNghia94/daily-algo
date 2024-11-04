package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{0, 4, 0, 3, 2}
	k := 1
	r := algo.FindMaxAverage(s, k)
	fmt.Println(r)
}
