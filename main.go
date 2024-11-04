package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	s := []string{"-60", "D", "-36", "30", "13", "C", "C", "-33", "53", "79"}
	r := algo.CalPoints(s)
	fmt.Println(r)
}
