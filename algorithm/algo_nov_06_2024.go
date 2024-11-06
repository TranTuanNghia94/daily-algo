package algo

import (
	"math/bits"
	"slices"
	"sort"
)

// https://leetcode.com/problems/find-if-array-can-be-sorted/description/?envType=daily-question&envId=2024-11-06
func CanSortArray(nums []int) bool {
	sort.SliceStable(nums, func(i, j int) bool {
		return bits.OnesCount(uint(nums[i])) == bits.OnesCount(uint(nums[j])) && nums[i] < nums[j]
	})

	return slices.IsSorted(nums)
}
