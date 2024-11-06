package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	a := "1111001110001111100001101"
	b := "101111001100000011011110011"
	// c := "10101"
	r := algo.AddBinary(a, b)
	fmt.Println("\n", r)
}
