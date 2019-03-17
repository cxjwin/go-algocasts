package leetcode581

import "testing"

func TestFindUnsortedSubarray(t *testing.T) {
	type testFunc func([]int) int
	testBody := func(f testFunc, t *testing.T) {
		nums := []int{2, 6, 4, 8, 10, 9, 15}
		res := f(nums)
		if res != 5 {
			t.Error("[2, 6, 4, 8, 10, 9, 15] -> 5")
		}

		nums = []int{2, 1}
		res = f(nums)
		if res != 2 {
			t.Error("[2, 1] -> 2")
		}

		nums = []int{1, 2, 3, 4}
		res = f(nums)
		if res != 0 {
			t.Error("[1,2,3,4] -> 0")
		}

		nums = []int{2, 6, 4, 8, 10, 9, 15}
		res = f(nums)
		if res != 5 {
			t.Error("[2,6,4,8,10,9,15] -> 5")
		}
	}

	// testBody(findUnsortedSubarray, t)
	// testBody(findUnsortedSubarrayTwoPass, t)
	testBody(findUnsortedSubarrayOnePass, t)
}
