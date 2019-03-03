package leetcode131

// https://leetcode.com/problems/palindrome-partitioning/

func _partition(s string, start int, d [][]bool, elems *[]string, res *[][]string) {
	if start >= len(s) {
		temp := make([]string, len(*elems))
		copy(temp, *elems)
		*res = append(*res, temp)
	} else {
		for end := start; end < len(s); end++ {
			if d[start][end] {
				*elems = append(*elems, s[start:end+1])
				_partition(s, end+1, d, elems, res)
				*elems = (*elems)[:len(*elems)-1] // T:O(1)
			}
		}
	}
}

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}

	res := &[][]string{}
	elems := &[]string{}

	n := len(s)
	d := make([][]bool, n)
	for i := 0; i < n; i++ {
		d[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if i == j {
				d[i][j] = true
			} else if i+1 == j {
				d[i][j] = (s[i] == s[j])
			} else {
				d[i][j] = (s[i] == s[j] && d[i+1][j-1])
			}
		}
	}

	_partition(s, 0, d, elems, res)

	return *res
}
