package leetcode62

import "testing"

func TestUniquePaths(t *testing.T) {
	if uniquePaths(3, 2) != 3 {
		t.Error("3 x 2 => 3")
	}
	if uniquePaths(7, 3) != 28 {
		t.Error("7 x 3 => 28")
	}
	if uniquePaths(2, 4) != 4 {
		t.Error("2 x 4 => 4")
	}
}
