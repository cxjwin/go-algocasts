package leetcode238

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	type testFunc func([]int) []int
	testBody := func(f testFunc, t *testing.T) {
		nums := []int{1, 4, 2, 8}
		if !reflect.DeepEqual(f(nums), []int{64, 16, 32, 8}) {
			t.Error("1, 4, 2, 8 => 64, 16, 32, 8")
		}
	}
	testBody(productExceptSelf, t)
	testBody(productExceptSelfO1, t)
	testBody(productExceptSelfNotChangeO1, t)
}
