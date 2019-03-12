package leetcode344

// https://leetcode.com/problems/reverse-string/

func reverseString(s []byte) {
	if s == nil || len(s) == 0 {
		return
	}

	n := len(s)
	mid := n / 2
	for i := 0; i < mid; i++ {
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}
