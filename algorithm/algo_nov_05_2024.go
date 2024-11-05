package algo

// https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/description/?envType=study-plan-v2&envId=leetcode-75
func MaxVowels(s string, k int) int {
	m := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	count := 0

	for i := 0; i < k; i++ {
		if m[rune(s[i])] {
			count++
		}
	}

	maxNum := count

	for i := k; i < len(s); i++ {
		if m[rune(s[i])] {
			count++
		}

		if m[rune(s[i-k])] {
			count--
		}

		if count > maxNum {
			maxNum = count
		}
	}

	return maxNum
}

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
