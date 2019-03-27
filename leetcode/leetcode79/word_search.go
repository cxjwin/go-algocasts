package leetcode79

// https://leetcode.com/problems/word-search/

func _exist(board [][]byte, visit [][]bool, i int, j int, word string, idx int) bool {
	if idx == len(word) {
		return true
	}
	m, n := len(board), len(board[0])
	if i < 0 || i >= m || j < 0 || j >= n || visit[i][j] || board[i][j] != word[idx] {
		return false
	}

	visit[i][j] = true

	existed := _exist(board, visit, i-1, j, word, idx+1) ||
		_exist(board, visit, i+1, j, word, idx+1) ||
		_exist(board, visit, i, j-1, word, idx+1) ||
		_exist(board, visit, i, j+1, word, idx+1)

	visit[i][j] = false

	return existed
}

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	if len(word) == 0 {
		return false
	}

	m, n := len(board), len(board[0])

	visit := make([][]bool, m)
	for i := 0; i < m; i++ {
		visit[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if _exist(board, visit, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}
