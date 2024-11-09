package algo

// https://leetcode.com/problems/minimum-array-end/description/?envType=daily-question&envId=2024-11-09
func MinEnd(n int, x int) int64 {
	shifts := make([]int16, 0)
	cur := int64(x)

	calc := int64(n - 1)

	for i := 0; i < 32; i++ {
		if ((1 << i) & x) == 0 {
			shifts = append(shifts, int16(i))
		}
	}

	for i := 32; i < 64; i++ {
		shifts = append(shifts, int16(i))
	}

	for i := 0; calc > 0; i++ {
		cur += (calc & 1) << shifts[i]

		calc >>= 1
	}

	return cur
}
