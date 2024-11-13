package algo

import (
	"sort"
)

// https://leetcode.com/problems/count-the-number-of-fair-pairs/description/?envType=daily-question&envId=2024-11-13
func CountFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	if n < 2 {
		return 0
	}

	sort.Ints(nums)
	var count int64

	for i := 0; i < n-1; i++ {
		left := searchLeft(nums, i+1, n-1, lower-nums[i])
		right := searchRight(nums, i+1, n-1, upper-nums[i])

		if left <= right {
			count += int64(right - left + 1)
		}
	}

	return count
}

func searchLeft(nums []int, left, right, target int) int {
	result := right + 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}

func searchRight(nums []int, left, right, target int) int {
	result := left - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			result = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}
