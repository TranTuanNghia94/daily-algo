package algo

// https://leetcode.com/problems/maximum-sum-of-distinct-subarrays-with-length-k/description/?envType=daily-question&envId=2024-11-19
func MaximumSubarraySum(nums []int, k int) int64 {
	if len(nums) < k {
		return 0
	}

	tmpNum := int64(0)
	maxSum := int64(0)
	m := make(map[int]int)

	for i := 0; i < k; i++ {
		tmpNum += int64(nums[i])
		m[nums[i]]++
	}
	if len(m) == k {
		maxSum = tmpNum
	}

	for i := k; i < len(nums); i++ {
		prev := nums[i-k]
		m[prev]--
		if m[prev] == 0 {
			delete(m, prev)
		}
		tmpNum -= int64(prev)

		next := nums[i]
		m[next]++
		tmpNum += int64(next)

		if len(m) == k {
			if tmpNum > maxSum {
				maxSum = tmpNum
			}
		}
	}

	return maxSum
}
