package leetcode493

import "testing"

func TestReversePairs(t *testing.T) {
	nums := []int{1, 3, 2, 3, 1}
	if reversePairsBruteForce(nums) != 2 {
		t.Error("a : 1,3,2,3,1 -> 2")
	}

	if reversePairs(nums) != 2 {
		t.Error("b : 1,3,2,3,1 -> 2")
	}
}
