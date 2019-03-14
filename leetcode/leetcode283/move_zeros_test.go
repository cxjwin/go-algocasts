package leetcode283

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	type testFunc func([]int)

	testBody := func(f testFunc, t *testing.T) {
		nums := []int{0, 1, 0, 3, 12}
		f(nums)
		if !reflect.DeepEqual(nums, []int{1, 3, 12, 0, 0}) {
			t.Error("[1,3,12,0,0]")
		}
	}

	testBody(moveZeroes, t)
	testBody(moveZerosAssign, t)
	testBody(moveZerosSwap, t)
}
