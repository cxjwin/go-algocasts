package leetcode207

import "testing"

func TestCanFinish(t *testing.T) {
	type testFunc func(int, [][]int) bool

	testBody := func(f testFunc, t *testing.T) {
		case1 := [][]int{
			{1, 0},
			{3, 0},
			{3, 1},
			{2, 1},
			{2, 3},
			{4, 2},
			{4, 3},
		}
		if !f(5, case1) {
			t.Error("case1 can finish")
		}

		case2 := [][]int{
			{1, 0},
		}
		if !f(2, case2) {
			t.Error("case2 can finish")
		}

		case3 := [][]int{
			{1, 0},
			{0, 1},
		}
		if f(2, case3) {
			t.Error("case3 can't finish")
		}
	}

	testBody(canFinish, t)
	testBody(canFinishTopoSort, t)
}
