package algo

import (
	"math"
)

// 2,4,9,3

func Decrypt(code []int, k int) []int {
	if k == 0 {
		return make([]int, len(code))
	}

	arr := make([]int, 0)
	n := len(code)

	for i := 0; i < n; i++ {
		sum := 0
		if k > 0 {
			for j := 1; j <= k; j++ {
				if i+j >= n {
					sum += code[i+j-n]
				} else {
					sum += code[i+j]
				}
			}

		} else {
			for j := 1; j <= -k; j++ {
				if i-j < 0 {
					num := math.Abs(float64(i - j))
					sum += code[n-int(num)]
				} else {
					sum += code[i-j]
				}
			}
		}

		arr = append(arr, sum)
	}

	return arr
}
