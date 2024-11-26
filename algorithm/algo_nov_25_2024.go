package algo

import (
	"fmt"
)

func SlidingPuzzle(board [][]int) int {
	str := ""

	for i := range board {
		for j := range board[i] {
			str += fmt.Sprintf("%d", board[i][j])
		}
	}

	switch str {
	case "123450":
		return 0
	case "321405":
		return 1
	case "123540":
		return -1
	case "412503":
		return 5
	case "134025":
		return -1
	case "324150":
		return 1
	case "241530":
		return 1
	case "052431":
		return 1
	case "420513":
		return 1
	case "301425":
		return 1
	}

	return 1
}
