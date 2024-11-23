package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	// s := []int{1, 5, 4, 2, 9, 9, 9}
	// k := 3
	s := [][]byte{{'*'}}
	r := algo.RotateTheBox(s)
	fmt.Println("\n", r)
}
