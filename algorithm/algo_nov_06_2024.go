package algo

import (
	"math/bits"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func LongestPalindrome(s string) string {
	start := 0
	end := len(s) - 1

	for ; start < end; start, end = start+1, end-1 {
	}

	return s[start : end+1]
}

// https://leetcode.com/problems/fibonacci-number/description/?envType=study-plan-v2&envId=dynamic-programming
func Fib(n int) int {
	a, b := 0, 1

	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return a
}

// https://leetcode.com/problems/add-binary/description/?envType=study-plan-v2&envId=programming-skills
func AddBinary(a string, b string) string {
	result := []string{}
	i := len(a) - 1
	j := len(b) - 1
	carry := 0

	for i >= 0 || j >= 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i]) - int('0')
		}

		if j >= 0 {
			sum += int(b[j]) - int('0')
		}

		result = append(result, strconv.Itoa(sum%2))

		if sum > 1 {
			carry = 1
		} else {
			carry = 0
		}
		i--
		j--
	}

	if carry != 0 {
		result = append(result, "1")
	}
	slices.Reverse(result)
	return strings.Join(result, "")
}

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
