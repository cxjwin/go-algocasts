package leetcode221

import "testing"

func TestMaximalSquare(t *testing.T) {
	input := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	res := maximalSquare(input)

	if res != 4 {
		t.Error("Output: 4")
	}
}
