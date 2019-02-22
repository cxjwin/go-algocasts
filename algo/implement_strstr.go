package algo

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 && len(needle) == 0 {
		return 0
	}

	if len(haystack) != 0 && len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 && len(needle) != 0 {
		return -1
	}

	n1 := len(haystack)
	n2 := len(needle)

	for i := 0; i <= n1-n2; i++ {
		j := 0
		k := i
		for j < n2 && k < n1 && needle[j] == haystack[k] {
			j++
			k++
		}
		if j == n2 {
			return i
		}
	}

	return -1
}
