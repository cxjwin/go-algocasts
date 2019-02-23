package algo

// https://leetcode.com/problems/edit-distance/

func min(a int, b int, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

func editDistance(s string, t string) int {
	m := len(s) + 1
	n := len(t) + 1

	// m*n matrix
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		d[i][0] = i
	}
	for j := 0; j < n; j++ {
		d[0][j] = j
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if s[i-1] == t[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j-1], d[i-1][j], d[i][j-1]) + 1
			}
		}
	}

	return d[m-1][n-1]
}

func editDistance1dArray(s string, t string) int {
	m := len(s) + 1
	n := len(t) + 1

	d := make([]int, n)

	for i := 0; i < n; i++ {
		d[i] = i
	}

	for i := 1; i < m; i++ {
		pre := d[0]
		d[0] = i
		for j := 1; j < n; j++ {
			temp := d[j]
			if s[i-1] == t[j-1] {
				d[j] = pre
			} else {
				d[j] = min(pre, d[j], d[j-1]) + 1
			}
			pre = temp
		}
	}

	return d[n-1]
}
