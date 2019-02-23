package algo

// leetcode : https://leetcode.com/problems/valid-palindrome/

// 是否是字母或数字
func isAlphanumeric(r byte) bool {
	switch {
	case 'a' <= r && r <= 'z',
		'A' <= r && r <= 'Z',
		'0' <= r && r <= '9':
		return true
	default:
		return false
	}
}

// 判断两个字符串是否相等(忽略大小写)
func isEqualIgnoreCase(a byte, b byte) bool {
	if 'A' <= a && a <= 'Z' {
		a += 32
	}
	if 'A' <= b && b <= 'Z' {
		b += 32
	}

	return a == b
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		for i < j && !isAlphanumeric(s[i]) {
			i++
		}
		for i < j && !isAlphanumeric(s[j]) {
			j--
		}
		if !isEqualIgnoreCase(s[i], s[j]) {
			return false
		}
	}

	return true
}
