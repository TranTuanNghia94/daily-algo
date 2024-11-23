package algo

// https://leetcode.com/problems/rotating-the-box/description/?envType=daily-question&envId=2024-11-23
func RotateTheBox(box [][]byte) [][]byte {
	if len(box) == 0 || len(box[0]) == 0 {
		return nil
	}

	rows, cols := len(box), len(box[0])
	rotated := make([][]byte, cols)
	for i := range rotated {
		rotated[i] = make([]byte, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = box[i][j]
		}
	}

	for col := 0; col < rows; col++ {
		bottom := cols - 1

		for cur := cols - 1; cur >= 0; cur-- {
			switch rotated[cur][col] {
			case '#':
				if cur != bottom {
					rotated[bottom][col] = '#'
					rotated[cur][col] = '.'
				}
				bottom--
			case '*':
				bottom = cur - 1
			}
		}
	}

	return rotated
}
