package leetcode448

import (
	"reflect"
	"testing"
)

func TestFindDisappearedNumbers(t *testing.T) {
	type testFunc func([]int) []int

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
		res := f(nums)
		if !reflect.DeepEqual(res, []int{5, 6}) {
			t.Error("[4,3,2,7,8,2,3,1] -> [5,6]")
		}
	}

	testBody(findDisappearedNumbers, t)
	testBody(findDisappearedNumbersO1, t)
}
