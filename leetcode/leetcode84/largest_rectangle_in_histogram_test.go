package leetcode84

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	type testFunc func([]int) int

	testBody := func(f testFunc, t *testing.T) {
		heights := []int{2, 1, 5, 6, 2, 3}
		res := f(heights)
		if res != 10 {
			t.Error("1 : Output: 10")
		}
	}

	testBody(largestRectangleAreaExpand, t)
	testBody(largestRectangleArea, t)
}
