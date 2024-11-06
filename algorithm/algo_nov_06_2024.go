package algo

import (
	"math/bits"
	"slices"
	"sort"
)

// https://leetcode.com/problems/monotonic-array/description/?envType=study-plan-v2&envId=programming-skills
func IsMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
		if nums[0] < nums[len(nums)-1] {
			if nums[i+1] < nums[i] || nums[j-1] > nums[j] {
				return false
			}
		} else if nums[0] >= nums[len(nums)-1] {
			if nums[i+1] > nums[i] || nums[j-1] < nums[j] {
				return false
			}
		}
	}

	return true
}

// https://leetcode.com/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-11-06
func CanSortArray(nums []int) bool {
	sort.SliceStable(nums, func(i, j int) bool {
		return bits.OnesCount(uint(nums[i])) == bits.OnesCount(uint(nums[j])) && nums[i] < nums[j]
	})

	return slices.IsSorted(nums)
}
