package leetcode63

// https://leetcode.com/problems/unique-paths-ii/

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}

	n := len(obstacleGrid)
	m := len(obstacleGrid[0])

	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, m)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		d[0][i] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[j][0] == 1 {
			break
		}
		d[j][0] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				d[i][j] = 0
			} else {
				d[i][j] = d[i-1][j] + d[i][j-1]
			}
		}
	}

	return d[n-1][m-1]
}
