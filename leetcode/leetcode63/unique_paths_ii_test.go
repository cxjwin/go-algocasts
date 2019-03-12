package leetcode63

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles(matrix)
	if res != 2 {
		t.Error("1 - 2")
	}

	matrix = [][]int{
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	res = uniquePathsWithObstacles(matrix)
	if res != 2 {
		t.Error("2 - 2")
	}
}
