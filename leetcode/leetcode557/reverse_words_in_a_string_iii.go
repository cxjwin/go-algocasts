package leetcode557

// https://leetcode.com/problems/reverse-words-in-a-string-iii/

func reverse(s []rune, start int, end int) {
	if start >= end {
		return
	}
	i, j := start, end
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}

	arr := make([]rune, len(s))
	start, end := 0, 0
	for i, c := range s {
		arr[i] = c
		if c == ' ' {
			end = i - 1
			reverse(arr, start, end)
			start = i + 1
		}
	}
	if start < len(s)-1 {
		reverse(arr, start, len(s)-1)
	}
	return string(arr)
}
