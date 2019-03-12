package leetcode152

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	res := maxProduct(nums)
	if res != 6 {
		t.Error("res is 6")
	}
	nums = []int{-2, 3, -4}
	res = maxProduct(nums)
	if res != 24 {
		t.Error("res is 24")
	}
}
