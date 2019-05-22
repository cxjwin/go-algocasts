package leetcode85

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleArea(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	n := len(heights)

	stack := make([]int, n+1)
	maxA := 0
	top := -1

	for r := 0; r <= n; r++ {
		h := 0
		if r != n {
			h = heights[r]
		}

		for top != -1 && h < heights[stack[top]] {
			idx := stack[top]
			top--
			l := -1
			if top != -1 {
				l = stack[top]
			}
			maxA = max(maxA, heights[idx]*(r-l-1))
		}

		top++
		stack[top] = r
	}

	return maxA
}

func maximalRectangle(matrix [][]byte) int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	d := make([]int, n)

	for j := 0; j < n; j++ {
		d[j] = 0
		if matrix[0][j] == '1' {
			d[j] = 1
		}
	}

	maxA := largestRectangleArea(d)

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				d[j]++
			} else {
				d[j] = 0
			}
		}
		maxA = max(maxA, largestRectangleArea(d))
	}

	return maxA
}
