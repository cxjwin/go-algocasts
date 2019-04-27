package leetcode45

import "testing"

func TestJump(t *testing.T) {
	type testFunc func([]int) int
	testBody := func(f testFunc, t *testing.T) {
		arr := []int{2, 3, 1, 1, 4}
		if f(arr) != 2 {
			t.Error("0 -> 1 -> end")
		}
		arr = []int{2, 1, 0, 4}
		if f(arr) != -1 {
			t.Error("can't jump to end")
		}
	}
	testBody(jump, t)
	testBody(jumpO1, t)
}
