package leetcode31

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	nums := []int{2, 1, 8, 4, 2, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, []int{2, 2, 1, 1, 4, 8}) {
		t.Error("next permutation is [2, 2, 1, 1, 4, 8]")
	}

	nums = []int{4, 2, 1}
	nextPermutation(nums)
	if !reflect.DeepEqual(nums, []int{1, 2, 4}) {
		t.Error("next permutation is [1, 2, 4]")
	}
}
