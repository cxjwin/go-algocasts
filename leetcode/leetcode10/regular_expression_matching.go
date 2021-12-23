package leetcode10

// https://leetcode.com/problems/regular-expression-matching/submissions/

func isMatch(s string, p string) bool {
	if p == "*" {
		return false
	}

	m, n := len(s)+1, len(p)+1
	d := make([][]bool, m)
	for i := 0; i < m; i++ {
		d[i] = make([]bool, n)
	}
	d[0][0] = true
	for j := 0; j < len(p); j++ {
		if p[j] == '*' && d[0][j-1] {
			d[0][j+1] = true
		}
	}

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if p[j] == '.' {
				d[i+1][j+1] = d[i][j]
			}
			if p[j] == s[i] {
				d[i+1][j+1] = d[i][j]
			}
			if p[j] == '*' {
				if p[j-1] == s[i] || p[j-1] == '.' {
					d[i+1][j+1] = d[i+1][j] || d[i][j+1] || d[i+1][j-1]
				} else {
					d[i+1][j+1] = d[i+1][j-1]
				}
			}
		}
	}

	return d[m-1][n-1]
}
