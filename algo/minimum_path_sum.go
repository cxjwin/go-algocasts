package algo

import "github.com/cxjwin/go-algocasts/utils"

// https://leetcode.com/problems/minimum-path-sum/

func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	d := make([][]int, m)
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
	}

	d[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		d[i][0] = d[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		d[0][j] = d[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = utils.IntMin(d[i-1][j], d[i][j-1]) + grid[i][j]
		}
	}

	return d[m-1][n-1]
}

func minPathSum1dArray(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m := len(grid)
	n := len(grid[0])

	d := make([]int, n)

	d[0] = grid[0][0]
	for j := 1; j < n; j++ {
		d[j] = d[j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		d[0] += grid[i][0]
		for j := 1; j < n; j++ {
			d[j] = utils.IntMin(d[j-1], d[j]) + grid[i][j]
		}
	}

	return d[n-1]
}
