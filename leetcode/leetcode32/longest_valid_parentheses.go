package leetcode32

// https://leetcode.com/problems/longest-valid-parentheses/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	stack := make([]int, n+1)
	maxL := 0
	top := 0
	stack[top] = -1

	for i := 0; i < n; i++ {
		// valid parentheses
		if stack[top] != -1 && s[stack[top]] == '(' && s[i] == ')' {
			top--
			maxL = max(maxL, i-stack[top])
		} else {
			top++
			stack[top] = i
		}
	}

	return maxL
}

func longestValidParenthesesDP(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	d := make([]int, n)
	left := 0
	maxL := 0

	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else if left > 0 {
			d[i] = d[i-1] + 2
			if i-d[i] >= 0 {
				d[i] += d[i-d[i]]
			}
			maxL = max(maxL, d[i])
			left--
		}
	}

	return maxL
}
