package leetcode647

func isPalindrome(s string, start int, end int) bool {
	if start >= end {
		return true
	}

	for i, j := start, end; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func count(s string, end int) int {
	count := 1
	for i := 0; i < end; i++ {
		if isPalindrome(s, i, end) {
			count++
		}
	}
	return count
}

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	n := len(s)
	d := make([]int, n)
	d[0] = 1

	for i := 1; i < n; i++ {
		d[i] = d[i-1] + count(s, i)
	}

	return d[n-1]
}
