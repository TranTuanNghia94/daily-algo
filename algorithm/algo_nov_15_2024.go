package algo

// https://leetcode.com/problems/shortest-subarray-to-be-removed-to-make-array-sorted/description/?envType=daily-question&envId=2024-11-15
func FindLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left, right := 0, n-1

	// Find the longest non-decreasing prefix
	for left+1 < n && arr[left] <= arr[left+1] {
		left++
	}

	// If the entire array is already sorted
	if left == n-1 {
		return 0
	}

	// Find the longest non-decreasing suffix
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}

	// Calculate the minimum length of subarray to remove
	ans := min(n-1-left, right)

	// Merge left and right subarrays to find the shortest subarray to remove
	for l := 0; l <= left; l++ {
		for right < n && arr[right] < arr[l] {
			right++
		}
		ans = min(ans, right-l-1)
	}

	return ans
}
