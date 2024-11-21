package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	// s := []int{1, 5, 4, 2, 9, 9, 9}
	// k := 3
	m := 4
	n := 6
	guards := [][]int{{0, 0}, {1, 1}, {2, 3}}
	walls := [][]int{{0, 1}, {2, 2}, {1, 2}}
	r := algo.CountUnguarded(m, n, guards, walls)
	fmt.Println("\n", r)
}
