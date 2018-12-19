package algo

func countPalindromicSubstringsDynamicProgramming(str string) int {
	n := len(str)
	if n == 0 {
		return 0
	}

	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}

	count := 0
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				d[i][j] = true
			} else if i+1 == j {
				d[i][j] = str[i] == str[j]
			} else {
				d[i][j] = str[i] == str[j] && d[i+1][j-1]
			}

			if d[i][j] {
				count++
			}
		}
	}

	return count
}

func expand(str string, left int, right int) int {
	count := 0
	for left >= 0 && right < len(str) {
		if str[left] == str[right] {
			count++
			left--
			right++
		} else {
			break
		}
	}
	return count
}

func countPalindromicSubstringsExpand(str string) int {
	length := len(str)
	if length == 0 {
		return 0
	}

	count := 0
	for i := 0; i < length; i++ {
		count += expand(str, i, i)
		count += expand(str, i, i+1)
	}
	return count
}
