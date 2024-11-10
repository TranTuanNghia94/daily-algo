package algo

// https://leetcode.com/problems/shortest-subarray-with-or-at-least-k-ii/?envType=daily-question&envId=2024-11-10
func MinimumSubarrayLength(nums []int, k int) int {
	if k == 0 {
		return 1
	}
	a := nums[0]
	for i := 1; i < len(nums); i++ {
		a |= nums[i]
	}
	if a < k {
		return -1
	}
	counter := [30]int{}
	minVal := len(nums)
	l := 0
	for r := 0; r < len(nums); r++ {
		for i := 0; i < 30; i++ {
			if (nums[r] & (1 << i)) > 0 {
				counter[i]++
			}
		}
		for l < r && calc(counter) >= k {
			minVal = min(minVal, r-l+1)
			for i := 0; i < 30; i++ {
				if (nums[l] & (1 << i)) > 0 {
					counter[i]--
				}
			}
			l++
		}
		if calc(counter) >= k {
			minVal = min(minVal, r-l+1)
		}
	}
	return minVal
}

func calc(counter [30]int) int {
	result := 0
	for i, n := range counter {
		if n > 0 {
			result += 1 << i
		}
	}
	return result
}
