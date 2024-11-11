package algo

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// https://leetcode.com/problems/prime-subtraction-operation/description/?envType=daily-question&envId=2024-11-11
func PrimeSubOperation(nums []int) bool {
	n := len(nums)

	for i := 0; i < n; i++ {
		for k := nums[i]; k >= 2; k-- {
			r := isPrime(k)
			if i == 0 && r && nums[i]-k > 0 {
				nums[i] = nums[i] - k
				break
			} else if r && nums[i]-k > 0 && nums[i]-k > nums[i-1] {
				nums[i] = nums[i] - k
				break
			}
		}
	}

	for i := 0; i < n-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
	}

	return true
}
