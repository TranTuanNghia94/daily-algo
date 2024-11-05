package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := "weallloveyou"
	k := 7
	r := algo.MaxVowels(s, k)
	fmt.Println(r)
}
