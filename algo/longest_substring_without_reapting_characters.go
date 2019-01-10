package algo

func lengthOfLongestSubstring2N(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	l := len(s)
	for i := 0; i < l; i++ {
		count := 0
		for j := 0; j < l; j++ {
			c, ok := m[s[j]]
			if ok {
				count = c
				break
			} else {
				m[s[j]] = 
			}
		}
		
		if maxLen < j - i {
			maxLen = j - i
		}
		count--
		m[s[i]] = count
	}

	return maxLen
}
