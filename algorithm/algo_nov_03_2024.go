package algo

import (
	"strings"
)

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		requests: make([]int, 0, 10000), // Pre-allocate capacity for better performance
	}
}

// https://leetcode.com/problems/number-of-recent-calls/description/?envType=study-plan-v2&envId=leetcode-75
func (a *RecentCounter) Ping(t int) int {
	// Add new timestamp
	a.requests = append(a.requests, t)

	// Binary search to find the first index within the 3000ms window
	windowStart := t - 3000
	left, right := 0, len(a.requests)-1
	firstValidIndex := right

	for left <= right {
		mid := (left + right) / 2
		if a.requests[mid] >= windowStart {
			firstValidIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	// Remove outdated timestamps and return count of valid ones
	a.requests = a.requests[firstValidIndex:]
	return len(a.requests)

}

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct *= nums[i]
	}

	return result
}

// https://leetcode.com/problems/rotate-string/description/?envType=daily-question&envId=2024-11-03
func RotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}

// https://leetcode.com/problems/container-with-most-water/description/?envType=study-plan-v2&envId=leetcode-75
func MaxArea(height []int) int {
	maxNum := 0

	i, j := 0, len(height)-1
	for i < j {
		if height[i] < height[j] {
			dist := j - i
			count := height[i] * dist
			if count > maxNum {
				maxNum = count
			}
			i++
		} else {
			dist := j - i
			count := height[j] * dist
			if count > maxNum {
				maxNum = count
			}
			j--
		}
	}

	return maxNum
}
