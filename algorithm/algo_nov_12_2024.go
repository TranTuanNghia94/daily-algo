package algo

import "sort"

func binarySearch(num int, arr [][]int) int {
	low, high := 0, len(arr)-1
	found := -1

	for low <= high {
		mid := (low + high) / 2
		if arr[mid][0] <= num {
			found = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return found
}

// https://leetcode.com/problems/most-beautiful-item-for-each-query/description/?envType=daily-question&envId=2024-11-12
func MaximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	n := len(items)
	maxBeauty := make([]int, n)
	maxBeauty[0] = items[0][1]

	for i := 1; i < n; i++ {
		if items[i][1] > maxBeauty[i-1] {
			maxBeauty[i] = items[i][1]
		} else {
			maxBeauty[i] = maxBeauty[i-1]
		}
	}

	results := make([]int, len(queries))
	memo := make(map[int]int)

	for i, val := range queries {
		if cachedValue, exists := memo[val]; exists {
			results[i] = cachedValue
			continue
		}

		j := binarySearch(val, items)

		if j == -1 {
			results[i] = 0
			memo[val] = 0
		} else {
			results[i] = maxBeauty[j]
			memo[val] = maxBeauty[j]
		}
	}

	return results
}
