package leetcode139

// https://leetcode.com/problems/word-break/

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	d := make([]bool, n+1)
	d[0] = true

	for i := 1; i <= n; i++ {
		for j := i - 1; j >= 0; j-- {
			if d[j] {
				sub := s[j:i]
				if _, ok := m[sub]; ok {
					d[i] = true
				}
			}
		}
	}

	return d[n]
}
