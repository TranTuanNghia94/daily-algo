package algo

// https://leetcode.com/problems/maximum-xor-for-each-query/description/?envType=daily-question&envId=2024-11-08
func GetMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	result := make([]int, n)
	maxNum := (1 << maximumBit) - 1

	currXor := 0

	for i := range nums {
		currXor ^= nums[i]
	}

	for i := 0; i < n; i++ {
		k := currXor ^ maxNum
		result[i] = k

		if i < n-1 {
			currXor ^= nums[n-1-i]
		}

	}

	return result
}
