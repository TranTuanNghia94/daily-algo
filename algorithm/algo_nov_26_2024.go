package algo

func FindChampion(n int, edges [][]int) int {
	arr := make([]int, n)

	for _, row := range edges {
		arr[row[1]]--
	}

	champ := -1
	countChamp := 0

	for i := range arr {
		if arr[i] == 0 {
			champ = i
			countChamp++
		}
	}

	if countChamp > 1 {
		return -1
	}

	return champ
}
