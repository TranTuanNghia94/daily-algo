package algo

func ShortestSubarray(nums []int, k int) int {
	n := len(nums)
	// Create prefix sum array
	prefix := make([]int64, n+1) // Using int64 to handle potential overflow
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + int64(nums[i])
	}

	// Initialize result with maximum possible value
	result := n + 1

	// Implement deque for sliding window
	// deque will store indices
	deque := make([]int, 0)

	for right := 0; right <= n; right++ {
		// While we have indices in deque and current prefix sum minus
		// prefix sum at deque's first index is at least k
		for len(deque) > 0 && prefix[right]-prefix[deque[0]] >= int64(k) {
			// Update result with minimum length found
			length := right - deque[0]
			if length < result {
				result = length
			}
			// Remove the leftmost index
			deque = deque[1:]
		}

		// Remove indices from right of deque while current prefix sum
		// is less than prefix sum at those indices
		// This maintains monotonic increasing property
		for len(deque) > 0 && prefix[right] <= prefix[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, right)
	}

	if result == n+1 {
		return -1
	}
	return result
}
