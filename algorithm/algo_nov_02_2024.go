package algo

import (
	"regexp"
	"strconv"
	"strings"
)

func GetSneakyNumbers(nums []int) []int {
	m := map[int]int{}
	r := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > 1 {
			r = append(r, nums[i])
		}

		if len(r) == 2 {
			break
		}
	}

	return r
}

func FindDifference(nums1 []int, nums2 []int) [][]int {
	arr1 := make([]int, 0)
	arr2 := make([]int, 0)
	map1 := map[int]int{}
	map2 := map[int]int{}

	for i := range nums1 {
		if _, ok := map1[nums1[i]]; !ok {
			map1[nums1[i]] = 1
		}
	}

	for i := range nums2 {
		if _, ok := map2[nums2[i]]; !ok {
			map2[nums2[i]] = 1
		}
	}

	for i := range map1 {
		if _, ok := map2[i]; !ok {
			arr1 = append(arr1, i)
		}
	}

	for i := range map2 {
		if _, ok := map1[i]; !ok {
			arr2 = append(arr2, i)
		}
	}

	if len(arr1) == 0 && len(arr2) == 0 {
		return [][]int{}
	}

	return [][]int{arr1, arr2}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Trap(height []int) int {
	maxLeft := 0
	maxRight := 0
	arrLeft := make([]int, 0)
	arrRight := make([]int, 0)

	count := 0

	for i, j := 0, len(height)-1; i < len(height); i, j = i+1, j-1 {
		if height[i] > maxLeft {
			maxLeft = height[i]
		}

		if height[j] > maxRight {
			maxRight = height[j]
		}

		arrLeft = append(arrLeft, maxLeft)
		arrRight = append([]int{maxRight}, arrRight...)
	}

	for i := range height {
		count += min(arrLeft[i], arrRight[i]) - height[i]
	}

	return count
}

func Compress(chars []byte) int {
	countWrite := 0
	size := len(chars)
	index := 0
	for index < size {
		count := 1
		for index+1 < size && chars[index] == chars[index+1] {
			index++
			count++
		}
		chars[countWrite] = chars[index]
		countWrite++
		if count > 1 {
			for _, c := range strconv.Itoa(count) {
				chars[countWrite] = byte(c)
				countWrite++
			}
		}
		index++
	}

	return countWrite
}

func IsCircularSentence(sentence string) bool {
	result := true
	arr := strings.Split(sentence, " ")

	for i := range arr {
		if i+1 < len(arr) {
			if arr[i][len(arr[i])-1] != arr[i+1][0] {
				result = false
				break
			}
		} else {
			if arr[i][len(arr[i])-1] != arr[0][0] {
				result = false
			}
		}
	}

	return result
}

func CapitalizeTitle(title string) string {
	title = strings.ToLower(title)

	words := strings.Split(title, " ")
	for i := range words {
		s := words[i]
		s = strings.ToUpper(s[:1]) + s[1:]
		words[i] = s
	}

	return strings.Join(words, " ")
}

func ReverseWords(s string) string {
	re := regexp.MustCompile(`\s+`)
	s = re.ReplaceAllString(s, " ")
	words := strings.Split(strings.TrimSpace(s), " ")
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return strings.Join(words, " ")
}

// [1,0,0,0,1,0,0]
func CanPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) < n || (len(flowerbed) == 1 && flowerbed[0] == 1 && n > 0) {
		return false
	}

	if len(flowerbed) == 1 || n == 0 {
		return true
	}

	i := 0
	for i < len(flowerbed) {
		if i == 0 && (flowerbed[i] == 0 && flowerbed[i+1] == 0) {
			n = n - 1
			i = i + 2
			if n == 0 {
				break
			}
			continue
		} else if i+1 < len(flowerbed) && (flowerbed[i] == 0 && flowerbed[i+1] == 0 && flowerbed[i-1] == 0) {
			n = n - 1
			i = i + 2
			if n == 0 {
				break
			}
		} else if i == len(flowerbed)-1 && (flowerbed[i] == 0 && flowerbed[i-1] == 0) {
			n = n - 1
			break
		} else {
			i = i + 1
		}
	}

	return n == 0
}

// [9,8,9]  + 1 -> [1,0,9,0]
func PlusOne(digits []int) []int {
	var r []int

	if digits[len(digits)-1]+1 <= 9 {
		r = append(digits[:len(digits)-1], digits[len(digits)-1]+1)
	} else {
		r = append(digits[:len(digits)-1], 0)

		for i := len(digits) - 2; i >= 0; i-- {
			if digits[i]+1 <= 9 {
				r[i] = digits[i] + 1
				break
			} else {
				r[i] = 0
			}
		}

		if r[0] == 0 {
			r = append([]int{1}, r...)
		}
	}

	return r
}

func RepeatedSubstringPattern(s string) bool {
	str := s + s

	return strings.Contains(str[1:len(str)-1], s)
}

type Vowels struct {
	value    string
	position int
}

// "IceCreAm" -> "AceCreIm"
func ReverseVowels(s string) string {
	arr := make([]Vowels, 0)
	m := map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1}

	newS := []rune(s)

	for i := 0; i < len(newS); i++ {
		if _, ok := m[strings.ToLower(string(newS[i]))]; ok {
			arr = append(arr, Vowels{string(newS[i]), i})
		}
	}

	for i := range arr {
		newS[arr[i].position] = rune(arr[len(arr)-i-1].value[0])
	}

	return string(newS)
}

func IsSubsequence(s string, t string) bool {
	result := false

	for i := 0; i < len(s); i++ {
		result = strings.IndexByte(t, byte(s[i])) >= 0
		if !result {
			break
		}

	}

	return result
}

func MyAtoi(s string) int {
	const (
		MaxInt32 = 1<<31 - 1     // 2147483647
		MinInt32 = -1 << 31      // -2147483648
		MaxDiv10 = MaxInt32 / 10 // Used for overflow checking
		MaxMod10 = MaxInt32 % 10 // Used for overflow checking
	)

	// Remove leading and trailing whitespace
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	// Handle sign
	start := 0
	negative := false
	switch s[0] {
	case '-':
		negative = true
		start = 1
	case '+':
		start = 1
	}

	// Process digits
	res := 0
	for i := start; i < len(s); i++ {
		// Break on first non-digit character
		if s[i] < '0' || s[i] > '9' {
			break
		}

		digit := int(s[i] - '0')

		// Check for overflow before multiplying
		if res > MaxDiv10 || (res == MaxDiv10 && digit > MaxMod10) {
			if negative {
				return MinInt32
			}
			return MaxInt32
		}

		res = res*10 + digit
	}

	if negative {
		return -res
	}
	return res
}

func RomanToInt(s string) int {
	m := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	total := 0
	n := len(s)

	for i := range s {
		if i+1 < n && m[s[i]] < m[s[i+1]] {
			total -= m[s[i]]
		} else {
			total += m[s[i]]
		}
	}

	return total
}

func MakeFancyString(s string) string {
	r := []rune(s)
	arr := make([]rune, 0)

	count := 1
	arr = append(arr, r[0])
	for i := 1; i < len(r); i++ {
		if r[i] == r[i-1] {
			count++
			if count < 3 {
				arr = append(arr, r[i])
			}
		}

		if r[i] != r[i-1] {
			count = 1
			arr = append(arr, r[i])
		}
	}

	return string(arr)

}

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxNum := candies[0]

	for _, v := range candies {
		if v > maxNum {
			maxNum = v
		}
	}

	result := make([]bool, len(candies))
	for i, v := range candies {
		result[i] = v+extraCandies >= maxNum
	}

	return result
}
