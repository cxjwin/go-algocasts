package leetcode300

import "testing"

func TestLengthOfLISBinarySearch(t *testing.T) {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	res := lengthOfLISBinarySearch(nums)
	if res != 4 {
		t.Error("length is 4")
	}
}
