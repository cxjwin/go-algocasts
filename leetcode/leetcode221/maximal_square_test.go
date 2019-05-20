package leetcode221

import "testing"

func TestMaximalSquare(t *testing.T) {
	type testFunc func([][]byte) int

	testBody := func(f testFunc, t *testing.T) {
		input := [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}
		res := f(input)

		if res != 4 {
			t.Error("Output: 4")
		}
	}

	testBody(maximalSquare, t)
	testBody(maximalSquareHistogram, t)
}
