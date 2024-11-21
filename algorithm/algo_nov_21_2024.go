package algo

import "fmt"

// m = 4, n = 6, guards = [[0,0],[1,1],[2,3]], walls = [[0,1],[2,2],[1,4]]
func CountUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]string, m)

	for i := 0; i < m; i++ {
		grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			grid[i][j] = "."
		}

	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = "W"
	}

	for _, guard := range guards {
		row := guard[0]
		col := guard[1]
		direc := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		grid[row][col] = "G"

		for _, d := range direc {
			dRow := d[0]
			dCol := d[1]
			nRow := row + dRow
			nCol := col + dCol

			for nRow >= 0 && nRow < m && nCol >= 0 && nCol < n {
				if grid[nRow][nCol] == "W" {
					break
				}
				grid[nRow][nCol] = "G"
				nRow += dRow
				nCol += dCol
			}

		}
	}

	count := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == "." {
				count++
			}
		}
	}

	fmt.Println(grid)

	return count
}
