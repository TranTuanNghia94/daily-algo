package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{1, 2, 3, 2, 5}
	r := algo.CanMakeArithmeticProgression(s)
	fmt.Println(r)
}
