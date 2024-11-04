package algo

import (
	"strconv"
)

// https://leetcode.com/problems/baseball-game/description/?envType=study-plan-v2&envId=programming-skills
func CalPoints(operations []string) int {
	sum := 0
	arr := make([]int, 0)
	for i := range operations {
		switch operations[i] {
		case "C":
			sum -= arr[len(arr)-1]
			arr = arr[:len(arr)-1]
		case "D":
			arr = append(arr, arr[len(arr)-1]*2)
			sum += arr[len(arr)-1]
		case "+":
			arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])
			sum += arr[len(arr)-1]
		default:
			num, err := strconv.Atoi(operations[i])
			if err != nil {
				return 0
			}
			arr = append(arr, num)
			sum += num
		}
	}

	return sum
}

// https://leetcode.com/problems/unique-number-of-occurrences/description/?envType=study-plan-v2&envId=leetcode-75
func UniqueOccurrences(arr []int) bool {
	var count = make(map[int]int)
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	var result = make(map[int]int)
	for _, value := range count {
		result[value]++
		if result[value] != 1 {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/maximum-average-subarray-i/?envType=study-plan-v2&envId=leetcode-75
func FindMaxAverage(nums []int, k int) float64 {
	maxAvgSum := 0.00000
	tmpSum := 0.00000
	for i := 0; i < k; i++ {
		maxAvgSum += float64(nums[i])
	}

	tmpSum = maxAvgSum

	for i := 1; i < len(nums); i++ {
		if i+k-1 < len(nums) {
			sum := tmpSum - float64(nums[i-1]) + float64(nums[i+k-1])
			tmpSum = sum
			if sum > maxAvgSum {
				maxAvgSum = sum
			}
		}

	}

	return float64(maxAvgSum / float64(k))
}

// https://leetcode.com/problems/string-compression-iii/description/?envType=daily-question&envId=2024-11-04
func CompressedString(word string) string {
	n := len(word)
	res := []byte{}
	i, j := 0, 0

	for i < n {
		c := word[i]
		for j = 1; j < 9; j++ {
			if i+j == n || c != word[i+j] {
				break
			}
		}
		res = append(res, byte(j)+'0')
		res = append(res, c)
		i += j
	}
	return string(res)
}
