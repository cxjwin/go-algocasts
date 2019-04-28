package leetcode120

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	res := minimumTotal(triangle)
	fmt.Println(res)
	if res != 11 {
		t.Error("The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).")
	}
}
