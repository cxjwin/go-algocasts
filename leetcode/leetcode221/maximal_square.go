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

func largestSquareInHistogram(heights []int) int {
	n := len(heights)

	st := make([]int, n+1)
	top := -1

	maxS := 0

	for r := 0; r <= n; r++ {
		h := 0
		if r != n {
			h = heights[r]
		}

		for top != -1 && h < heights[st[top]] {
			idx := st[top]
			top--

			l := -1
			if top != -1 {
				l = st[top]
			}

			width := r - l - 1
			len := min(width, heights[idx])

			maxS = max(maxS, len*len)
		}

		top++
		st[top] = r
	}

	return maxS
}

func maximalSquareHistogram(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	d := make([]int, n)

	maxS := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				d[j]++
			} else {
				d[j] = 0
			}
		}

		maxS = max(maxS, largestSquareInHistogram(d))
	}

	return maxS
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
