package leetcode42

import "testing"

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	sum := trap(height)
	if sum != 6 {
		t.Error("sum is 6")
	}
}

func TestTrapO1(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	sum := trapO1(height)
	if sum != 6 {
		t.Error("sum is 6")
	}
}
