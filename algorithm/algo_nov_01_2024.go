package algo

import (
	"fmt"
	"strconv"
)

func MergeAlternately(word1 string, word2 string) string {
	w1 := 0
	w2 := 0

	r := []rune{}

	for _, v := range word1 {
		r = append(r, v)
		if w2 < len(word2) {
			r = append(r, rune(word2[w2]))
			w2++
		}
		w1++
	}

	for w2 < len(word2) {
		r = append(r, rune(word2[w2]))
		w2++
	}

	return string(r)
}

func AddDigits(num int) int {
	if num <= 9 {
		return num
	} else if num%9 == 0 {
		return 9
	} else {
		return num % 9
	}
}

func SummaryRanges(nums []int) []string {
	var result []string

	if len(nums) == 0 {
		return result
	}

	start, end := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != end+1 {
			if start == end {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, end))
			}
			start = nums[i]
		}
		end = nums[i]
	}

	// Add the last range
	if start == end {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, end))
	}

	return result
}

func TwoSum(nums []int, target int) []int {
	pos := make(map[int]int, len(nums))
	for i, v := range nums {
		p2, ok := pos[target-v]
		if !ok {
			pos[v] = i
		} else {
			return []int{i, p2}
		}
	}

	panic("nonono")
}

func IsPowerOfTwo(n int) bool {
	switch {
	case n == 0:
		return true
	case n == 1:
		return true
	case n%2 != 0:
		return false
	}

	r := n / 2
	return IsPowerOfTwo(r)
}
