package leetcode34

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	type testFunc func([]int, int) []int

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		res := f(nums, 8)
		if !reflect.DeepEqual(res, []int{3, 4}) {
			t.Error("[5, 7, 7, 8, 8, 10] - 8 --> [3, 4]")
		}
		res = f(nums, 6)
		if !reflect.DeepEqual(res, []int{-1, -1}) {
			t.Error("[5, 7, 7, 8, 8, 10] - 6 --> [-1, -1]")
		}
	}

	// testBody(searchRange, t)
	testBody(searchRangeBinarySearch, t)
}
