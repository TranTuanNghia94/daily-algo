package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []int{2, 4, 9, 3}
	k := -2
	r := algo.Decrypt(s, k)
	fmt.Println("\n", r)
}
