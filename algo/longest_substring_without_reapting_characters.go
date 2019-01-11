package algo

func lengthOfLongestSubstring2N(s string) int {
	m := make(map[byte]int)
	maxLen := 0
	l := len(s)
	i := 0
	j := 0
	for ; i < l; i++ {
		for ; j < l; j++ {
			count, ok := m[s[j]]
			if ok && count > 0 {
				break
			} else {
				m[s[j]] = 1
			}
		}

		if j-i > maxLen {
			maxLen = j - i
		}

		m[s[i]]--
	}

	return maxLen
}

func lengthOfLongestSubstring1N(s string) int {
	indexMap := make(map[byte]int)
	maxLen := 0
	l := len(s)
	i := 0
	for j := 0; j < l; j++ {
		idx, ok := indexMap[s[j]]
		if ok {
			if idx+1 > i {
				i = idx + 1
			}
		}
		curLen := j - i + 1
		if curLen > maxLen {
			maxLen = curLen
		}
		indexMap[s[j]] = j
	}

	return maxLen
}
