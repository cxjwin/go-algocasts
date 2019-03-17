package leetcode387

// https://leetcode.com/problems/first-unique-character-in-a-string/

func firstUniqChar(s string) int {
	if len(s) == 0 {
		return -1
	}

	c := [26]int{}
	for _, v := range s {
		c[v-'a']++
	}
	for i, v := range s {
		if c[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func firstUniqCharOnePass(s string) int {
	if len(s) == 0 {
		return -1
	}

	c := [26]int{}
	pos := [26]int{}
	for i := range pos {
		pos[i] = -1
	}
	for i, v := range s {
		idx := v - 'a'
		c[idx]++
		if pos[idx] == -1 {
			pos[idx] = i
		}
	}

	minPos := 1<<31 - 1
	for i := 0; i < 26; i++ {
		if c[i] == 1 {
			minPos = min(minPos, pos[i])
		}
	}
	if minPos == 1<<31-1 {
		return -1
	}
	return minPos
}
