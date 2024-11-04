package algo

func CompressedString(word string) string {
	n := len(word)
	res := []byte{}
	i, j := 0, 0

	for i < n {
		c := word[i]
		for j = 1; j < 9; j++ {
			if i+j == n || c != word[i+j] {
				break
			}
		}
		res = append(res, byte(j)+'0')
		res = append(res, c)
		i += j
	}
	return string(res)
}
