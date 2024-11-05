package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := "0000"
	r := algo.MinChanges(s)
	fmt.Println(r)
}
