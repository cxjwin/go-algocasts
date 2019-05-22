package lintcode532

import "testing"

func TestReversePairs(t *testing.T) {
	nums := []int{2, 4, 1, 3, 5}
	if reversePairs(nums) != 3 {
		t.Error("2, 4, 1, 3, 5 -> 3")
	}
}
