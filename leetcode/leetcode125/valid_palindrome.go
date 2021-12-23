package leetcode125

// https://leetcode.com/problems/valid-palindrome/

func isAlphanumeric(ch uint8) bool {
	if ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9') {
		return true
	}
	return false
}

func isEqualIgnoreCase(c1 uint8, c2 uint8) bool {
	if 'A' <= c1 && c1 <= 'Z' {
		c1 += 32
	}
	if 'A' <= c2 && c2 <= 'Z' {
		c2 += 32
	}
	return c1 == c2
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// 这里需要注意，每一个循环都需要校验 i < j，不然会有越界的问题
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}
		if i < j && !isEqualIgnoreCase(s[i], s[j]) {
			return false
		}
	}
	return true
}
