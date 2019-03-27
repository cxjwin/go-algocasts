package leetcode200

import "testing"

func TestNumIslands(t *testing.T) {
	type testFunc func([][]byte) int

	testBody := func(f testFunc, t *testing.T) {
		case1 := [][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'1', '1', '1', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}
		res := f(case1)

		if res != 1 {
			t.Error("case1's result is 1")
		}

		case2 := [][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}

		res = f(case2)

		if res != 3 {
			t.Error("case2's result is 3")
		}
	}

	testBody(numIslands, t)
	testBody(numIslandsIterative, t)
}
