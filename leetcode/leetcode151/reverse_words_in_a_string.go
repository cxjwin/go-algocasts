package leetcode151

// https://leetcode.com/problems/reverse-words-in-a-string/

func reverse(s []byte, i int, j int) {
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

	arr := []byte(s)
	p, q, end := 0, 0, len(arr)-1

	for end >= 0 && arr[end] == ' ' {
		end--
	}

	for q <= end {
		start := p

		// skip space
		for q <= end && arr[q] == ' ' {
			q++
		}

		// copy single word
		for q <= end && arr[q] != ' ' {
			arr[p] = arr[q]
			p++
			q++
		}

		// reverse
		reverse(arr, start, p-1)

		// add one space
		if q <= end {
			arr[p] = ' '
			p++
		}
	}

	arr = arr[0:p]
	reverse(arr, 0, p-1)

	return string(arr)
}
