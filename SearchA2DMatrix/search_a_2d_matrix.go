package algo

func searchA2DMatrix(target int, matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{-1, -1}
	}

	if len(matrix[0]) == 0 {
		return []int{-1, -1}
	}

	m := len(matrix)
	n := len(matrix[0])
	i := 0
	j := n - 1
	for i < m && j >= 0 {
		arr := matrix[i]
		num := arr[j]
		if target < num {
			j--
		} else if target > num {
			i++
		} else {
			return []int{i, j}
		}
	}

	return []int{-1, -1}
}
