package leetcode64

import "math"

func _minPathSumRecursive(grid [][]int, i int, j int, curPath int, minPath *int) {
	curPath += grid[i][j]

	if i == len(grid)-1 && j == len(grid[i])-1 {
		if curPath < *minPath {
			*minPath = curPath
		}
		return
	}

	// 向下
	if i < len(grid)-1 {
		_minPathSumRecursive(grid, i+1, j, curPath, minPath)
	}
	// 向右
	if j < len(grid[i])-1 {
		_minPathSumRecursive(grid, i, j+1, curPath, minPath)
	}
}

// 暴力回溯法
func minPathSumRecursive(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	minPath := math.MaxInt32
	p := &minPath
	_minPathSumRecursive(grid, 0, 0, 0, p)
	return *p
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// 动态规划
func minPathSumDP(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	d := make([][]int, m)
	for i := range d {
		d[i] = make([]int, n)
	}
	d[0][0] = grid[0][0]
	// first column
	for i := 1; i < m; i++ {
		d[i][0] = d[i-1][0] + grid[i][0]
	}
	// first row
	for j := 1; j < n; j++ {
		d[0][j] = d[0][j-1] + grid[0][j]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			d[i][j] = min(d[i-1][j], d[i][j-1]) + grid[i][j]
		}
	}

	return d[m-1][n-1]
}
