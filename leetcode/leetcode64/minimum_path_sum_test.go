package leetcode64

import "testing"

func TestMinPathSum(t *testing.T) {
	type testFunc func([][]int) int
	testBody := func(f testFunc, t *testing.T) {
		var grid [][]int
		var res int

		grid = [][]int{
			{1, 3, 1},
			{1, 5, 1},
			{4, 2, 1},
		}
		res = f(grid)
		if res != 7 {
			t.Error("Because the path 1→3→1→1→1 minimizes the sum.")
		}

		grid = [][]int{
			{1, 2},
			{5, 6},
			{1, 1},
		}
		res = f(grid)
		if res != 8 {
			t.Error("Because the path 1→2→6→1 minimizes the sum.")
		}

		grid = [][]int{
			{1, 2, 3},
			{4, 5, 6},
		}
		res = f(grid)
		if res != 12 {
			t.Error("Because the path 1→2→3→6 minimizes the sum.")
		}
	}
	testBody(minPathSumRecursive, t)
	testBody(minPathSumDP, t)
}
