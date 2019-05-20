package leetcode84

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func largestRectangleAreaExpand(heights []int) int {
	if heights == nil || len(heights) == 0 {
		return 0
	}

	n := len(heights)
	maxA := 0
	for i := 0; i < n; i++ {
		l, r := i, i
		for l >= 0 && heights[l] >= heights[i] {
			l--
		}
		for r < n && heights[r] >= heights[i] {
			r++
		}
		maxA = max(maxA, heights[i]*(r-l-1))
	}

	return maxA
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
