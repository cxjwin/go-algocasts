package algo

func searchA2DMatrix(target int, matrix [][]int) (int, int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1, -1
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
			return i, j
		}
	}

	return -1, -1
}

func binarySearchA2DMatrix(target int, matrix [][]int) (int, int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1, -1
	}

	m := len(matrix)
	n := len(matrix[0])
	low := 0
	high := m*n - 1

	for low <= high {
		mid := low + (high-low)/2
		r, c := mid/n, mid%n
		if target == matrix[r][c] {
			return r, c
		}

		if target < matrix[r][c] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1, -1
}
