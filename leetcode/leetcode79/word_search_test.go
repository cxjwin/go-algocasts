package leetcode79

import "testing"

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	if !exist(board, "ABCCED") {
		t.Error("ABCCED => true")
	}

	if !exist(board, "SEE") {
		t.Error("SEE => true")
	}

	if exist(board, "ABCB") {
		t.Error("ABCB => false")
	}
}
