package leetcode85

import "testing"

func TestMaximalRectangle(t *testing.T) {
	var input [][]byte
	var res int

	input = [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	res = maximalRectangle(input)
	if res != 6 {
		t.Error("1 : output is 6")
	}

	input = [][]byte{
		{'1'},
	}

	res = maximalRectangle(input)
	if res != 1 {
		t.Error("2 : output is 1")
	}

	input = [][]byte{
		{'1', '1'},
	}

	res = maximalRectangle(input)
	if res != 2 {
		t.Error("3 : output is 2")
	}
}
