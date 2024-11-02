package algo

import (
	"strconv"
	"strings"
)

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func IsPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func IsAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapS, mapT := make(map[byte]int), make(map[byte]int)
	for i := range len(s) {
		mapS[s[i]]++
		mapT[t[i]]++
	}
	for k := range mapS {
		if mapS[k] != mapT[k] {
			return false
		}
	}
	return true
}

func CompressString(s string) string {
	if len(s) == 0 {
		return ""
	}
	var sb strings.Builder
	count := 1
	currentChar := s[0]

	for i := 1; i < len(s); i++ {
		if s[i] == currentChar {
			count++
		} else {
			sb.WriteByte(currentChar)
			if count > 1 {
				sb.WriteString(strconv.Itoa(count))
			}
			currentChar = s[i]
			count = 1
		}
	}

	// Handle the last group
	sb.WriteByte(currentChar)
	if count > 1 {
		sb.WriteString(strconv.Itoa(count))
	}

	return sb.String()
}
