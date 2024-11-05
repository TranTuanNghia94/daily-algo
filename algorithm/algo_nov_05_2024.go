package algo

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
