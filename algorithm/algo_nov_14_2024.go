package algo

// https://leetcode.com/problems/minimized-maximum-of-products-distributed-to-any-store/description/
func MinimizedMaximum(n int, quantities []int) int {
	left, right := 1, max(quantities)

	for left < right {
		mid := (left + right) / 2
		if canDistribute(quantities, n, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func canDistribute(quantities []int, n int, maxProducts int) bool {
	totalStores := 0
	for _, quantity := range quantities {
		totalStores += (quantity + maxProducts - 1) / maxProducts
	}
	return totalStores <= n
}

func max(arr []int) int {
	maxVal := arr[0]
	for _, v := range arr {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}
