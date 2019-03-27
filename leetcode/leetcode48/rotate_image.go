package leetcode48

// https://leetcode.com/problems/rotate-image/

func rotate(matrix [][]int) {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}
