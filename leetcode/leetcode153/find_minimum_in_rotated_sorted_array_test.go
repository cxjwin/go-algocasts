package leetcode153

import "testing"

func TestFindMin(t *testing.T) {
	type testFunc func([]int) int
	testBody := func(f testFunc, t *testing.T) {
		nums := []int{4, 5, 6, 7, 0, 1, 2}
		if f(nums) != 0 {
			t.Error("4, 5, 6, 7, 0, 1, 2 -> 0")
		}
	}
	testBody(findMin, t)
	testBody(findMinEarlyReturn, t)
}
