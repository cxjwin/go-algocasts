package leetcode221

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	maxL := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || matrix[i][j] == '0' {
				if matrix[i][j] == '1' {
					d[i][j] = 1
				} else {
					d[i][j] = 0
				}
			} else {
				d[i][j] = min(d[i-1][j-1], min(d[i][j-1], d[i-1][j])) + 1
			}
			maxL = max(maxL, d[i][j])
		}
	}

	return maxL * maxL
}
