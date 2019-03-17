package leetcode33

import "testing"

func TestSearch(t *testing.T) {
	nums := []int{5, 6, 7, 8, 1, 2, 3, 4}
	res := search(nums, 6)
	if res != 1 {
		t.Error("6 index is 1")
	}
	res = search(nums, 0)
	if res != -1 {
		t.Error("0 is not found")
	}
}
