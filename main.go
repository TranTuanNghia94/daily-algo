package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := "abcde"
	r := algo.CompressedString(s)
	fmt.Println(r)
}
