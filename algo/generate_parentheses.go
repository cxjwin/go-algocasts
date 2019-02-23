package algo

// https://leetcode.com/problems/generate-parentheses/

func generate(res *[]string, str string, left int, right int) {
	if left == 0 && right == 0 {
		temp := append((*res), str)
		// use slice pointer to cache result
		*res = temp
	}

	if left > 0 {
		generate(res, str+"(", left-1, right)
	}
	if right > left {
		generate(res, str+")", left, right-1)
	}
}

func generateParentheses(n int) []string {
	res := make([]string, 0)

	if n > 0 {
		temp := &res
		generate(temp, "", n, n)
		res = (*temp)
	}

	return res
}

func generateParenthesesDP(n int) []string {
	if n <= 0 {
		return make([]string, 0)
	}

	dp := make([][]string, n+1)
	for i := 0; i < n; i++ {
		dp[i] = make([]string, 0)
	}

	dp[0] = append(dp[0], "")

	for i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			for _, left := range dp[j] {
				for _, right := range dp[i-j-1] {
					dp[i] = append(dp[i], "("+left+")"+right)
				}
			}
		}
	}

	return dp[n]
}
