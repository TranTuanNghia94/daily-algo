package algo

// https://leetcode.com/problems/take-k-of-each-character-from-left-and-right/description/?envType=daily-question&envId=2024-11-20
func TakeCharacters(s string, k int) int {
	target := [3]int{}

	for i := range s {
		switch s[i] {
		case 'a':
			target[0]++
		case 'b':
			target[1]++
		case 'c':
			target[2]++
		}
	}

	for i, count := range target {
		if count-k < 0 {
			return -1
		}
		target[i] = count - k
	}

	window := [3]int{}
	out := len(s)
	l := 0

	for r := 0; r < len(s); r++ {
		switch s[r] {
		case 'a':
			window[0]++
		case 'b':
			window[1]++
		case 'c':
			window[2]++
		}

		for window[0] > target[0] || window[1] > target[1] || window[2] > target[2] {
			switch s[l] {
			case 'a':
				window[0]--
			case 'b':
				window[1]--
			case 'c':
				window[2]--
			}
			l++
		}

		out = min(out, len(s)-(r-l+1))
	}

	return out
}
