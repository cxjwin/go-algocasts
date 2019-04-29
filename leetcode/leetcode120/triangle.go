package leetcode120

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	if triangle == nil || len(triangle) == 0 {
		return 0
	}

	n := len(triangle)

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, i+1)
	}
	d[0][0] = triangle[0][0]
	for i := 1; i < n; i++ {
		d[i][0] = d[i-1][0] + triangle[i][0]
		d[i][i] = d[i-1][i-1] + triangle[i][i]
	}

	for i := 2; i < n; i++ {
		for j := 1; j < i; j++ {
			d[i][j] = min(d[i-1][j-1], d[i-1][j]) + triangle[i][j]
		}
	}

	minPath := d[n-1][0]
	for i := 1; i < n; i++ {
		if d[n-1][i] < minPath {
			minPath = d[n-1][i]
		}
	}
	return minPath
}
