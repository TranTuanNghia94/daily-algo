package main

import (
	"fmt"
	algo "hello/algorithm"
)

func main() {
	nums := []int{0, 3, 2, 1, 3, 2}
	r := algo.GetSneakyNumbers(nums)
	fmt.Println(r)
}
