package algo

import (
	"strings"
)

// https://leetcode.com/problems/flip-columns-for-maximum-number-of-equal-rows/description/?envType=daily-question&envId=2024-11-22
func MaxEqualRowsAfterFlips(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m := make(map[string]int)

	for _, row := range matrix {
		var str strings.Builder
		str.Grow(len(row))

		shouldFlip := row[0] == 1

		for _, bit := range row {
			if shouldFlip {
				str.WriteByte(byte('0' + (1 - bit)))
			} else {
				str.WriteByte(byte('0' + bit))
			}
		}

		m[str.String()]++
	}

	maxCount := 0
	for _, count := range m {
		if count > maxCount {
			maxCount = count
		}
	}

	return maxCount
}
