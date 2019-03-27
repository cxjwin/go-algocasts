package leetcode34

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{1, 2, 2, 4, 4, 8, 8}
	res := searchRange(nums, 2)
	if !reflect.DeepEqual(res, []int{1, 2}) {
		t.Error("2 int [1, 2, 2, 4, 4, 8, 8] is [1, 2]")
	}
}
