package leetcode200

// https://leetcode.com/problems/number-of-islands/

func dfs(grid [][]byte, visit [][]bool, i int, j int) {
	m, n := len(grid), len(grid[0])
	if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == '0' || visit[i][j] {
		return
	}
	visit[i][j] = true
	dfs(grid, visit, i-1, j)
	dfs(grid, visit, i+1, j)
	dfs(grid, visit, i, j-1)
	dfs(grid, visit, i, j+1)
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	num := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || visit[i][j] {
				continue
			}
			num++
			dfs(grid, visit, i, j)
		}
	}

	return num
}

type point struct {
	i int
	j int
}

func dfsIterative(grid [][]byte, visit [][]bool, i int, j int) {
	m, n := len(grid), len(grid[0])
	stack := []point{point{i: i, j: j}}
	for len(stack) != 0 {
		// pop
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// check
		if p.i < 0 || p.i >= m || p.j < 0 || p.j >= n || grid[p.i][p.j] == '0' || visit[p.i][p.j] {
			continue
		}

		visit[p.i][p.j] = true

		// push
		stack = append(stack, point{i: p.i - 1, j: p.j}, point{i: p.i + 1, j: p.j}, point{i: p.i, j: p.j - 1}, point{i: p.i, j: p.j + 1})
	}
}

func numIslandsIterative(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	num := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' || visit[i][j] {
				continue
			}
			num++
			dfsIterative(grid, visit, i, j)
		}
	}

	return num
}
