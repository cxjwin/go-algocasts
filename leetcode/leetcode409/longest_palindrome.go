package leetcode409

// https://leetcode.com/problems/longest-palindrome/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestPalindrome(s string) int {
	if len(s) == 0 {
		return 0
	}

	m := make(map[rune]int)
	for _, v := range s {
		cnt, ok := m[v]
		if ok {
			m[v] = cnt + 1
		} else {
			m[v] = 1
		}
	}

	odd := 0
	for _, v := range m {
		if v&1 == 1 {
			odd++
		}
	}

	unUsed := max(0, odd-1)

	return len(s) - unUsed
}
