package algo

import "github.com/cxjwin/go-algocasts/utils"

// https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindromicSubtringDP(s string) string {
	if len(s) == 0 {
		return s
	}

	n := len(s)
	// n * n matrix
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	start, maxLen := 0, 0

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = true
			} else if i+1 == j {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (s[i] == s[j] && dp[i+1][j-1])
			}

			if dp[i][j] && j-i+1 > maxLen {
				start = i
				maxLen = j - i + 1
			}
		}
	}

	return s[start : start+maxLen]
}

func expandLen(s string, left int, right int) int {
	i, j := left, right
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1
}

func longestPalindromicSubtringExpand(s string) string {
	if len(s) == 0 {
		return s
	}

	start, maxLen := 0, 0
	n := len(s)
	for i := 0; i < n; i++ {
		l1 := expandLen(s, i, i)
		l2 := expandLen(s, i, i+1)
		l := utils.IntMax(l1, l2)
		if l > maxLen {
			start = i - (l-1)/2
			maxLen = l
		}
	}

	return s[start : start+maxLen]
}
