package algo

// https://leetcode.com/problems/n-th-tribonacci-number/description/?envType=study-plan-v2&envId=dynamic-programming
func Tribonacci(n int) int {
	if n <= 3 {
		return n
	}

	a, b, c := 1, 1, 2

	for i := 3; i < n; i++ {
		a, b, c = b, c, a+b+c
	}

	return c
}

// https://leetcode.com/problems/climbing-stairs/solutions/?envType=study-plan-v2&envId=dynamic-programming
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}

	a, b := 1, 1

	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}

// https://leetcode.com/problems/largest-combination-with-bitwise-and-greater-than-zero/?envType=daily-question&envId=2024-11-07
func LargestCombination(candidates []int) int {
	bitCounts := make([]int, 32)
	max := 1
	for _, num := range candidates {
		for bit := 0; num > 0; bit++ {
			bitCounts[bit] += num & 1
			num >>= 1

			if bitCounts[bit] > max {
				max = bitCounts[bit]
			}
		}
	}

	return max
}
