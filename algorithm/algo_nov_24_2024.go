package algo

import (
	"math"
)

// https://leetcode.com/problems/maximum-matrix-sum/description/?envType=daily-question&envId=2024-11-24
func MaxMatrixSum(matrix [][]int) int64 {
	countNegative := 0
	minNumb := math.Abs(float64(matrix[0][0]))
	sum := 0

	for _, row := range matrix {
		for _, col := range row {
			if col < 0 {
				countNegative++
			}

			if math.Abs(float64(col)) < minNumb {
				minNumb = math.Abs(float64(col))
			}

			sum += int(math.Abs(float64(col)))
		}
	}

	if countNegative%2 == 0 {
		return int64(sum)
	}

	return int64(sum - (int(minNumb) * 2))
}
