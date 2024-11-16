package algo

// https://leetcode.com/problems/find-the-power-of-k-size-subarrays-i/description/
func ResultsArray(nums []int, k int) []int {
	n := len(nums)
	arr := make([]int, 0)

	for i := 0; i < n-k+1; i++ {

		flag := 0

		for j := i; j < i+k-1; j++ {
			if nums[j]+1 != nums[j+1] {
				flag = -1
				break
			}
		}

		if flag == -1 {
			arr = append(arr, -1)
		} else {
			arr = append(arr, nums[i+k-1])
		}

	}

	return arr
}
