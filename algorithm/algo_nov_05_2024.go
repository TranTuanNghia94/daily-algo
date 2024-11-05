package algo

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence/description/
func CanMakeArithmeticProgression(arr []int) bool {
	maxNum := arr[0]
	minNum := arr[0]
	m := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]]++

		if arr[i] < minNum {
			minNum = arr[i]
		}

		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}

	stepNum := (maxNum - minNum) / (len(arr) - 1)

	if len(m) == 1 && stepNum == 0 {
		return true
	}

	for k := range m {
		a := k - stepNum
		b := k + stepNum

		if (m[a] == 0 && m[b] == 0) || m[a] > 1 || m[b] > 1 {
			return false
		}
	}

	return true
}

// https://leetcode.com/problems/minimum-number-of-changes-to-make-binary-string-beautiful/?envType=daily-question&envId=2024-11-05
func MinChanges(s string) int {
	count := 0
	for i := 1; i < len(s); i += 2 {
		if s[i] != s[i-1] {
			count++
		}
	}
	return count
}
