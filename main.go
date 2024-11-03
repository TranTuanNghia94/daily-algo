package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := "bbbacddceeb"
	g := "ceebbbbacdd"
	r := algo.RotateString(s, g)
	fmt.Println(r)
}
