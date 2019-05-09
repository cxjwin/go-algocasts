package leetcode718

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLength(A []int, B []int) int {
	if A == nil || len(A) == 0 || B == nil || len(B) == 0 {
		return 0
	}

	m, n := len(A), len(B)

	d := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		d[i] = make([]int, n+1)
	}
	maxL := 0

	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 {
				d[i][j] = 0
			} else if A[i-1] == B[j-1] {
				d[i][j] = 1 + d[i-1][j-1]
				maxL = max(maxL, d[i][j])
			}
		}
	}

	return maxL
}
