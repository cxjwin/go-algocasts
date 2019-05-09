package leetcode718

import "testing"

func TestFindLength(t *testing.T) {
	A := []int{1, 2, 3, 2, 1}
	B := []int{3, 2, 1, 4, 7}
	res := findLength(A, B)
	if res != 3 {
		t.Error("The repeated subarray with maximum length is [3, 2, 1].")
	}
}
