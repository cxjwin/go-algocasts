package leetcode55

import "testing"

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	res := canJump(nums)
	if !res {
		t.Error("2,3,1,1,4 -> true")
	}
	nums = []int{3, 2, 1, 0, 4}
	res = canJump(nums)
	if res {
		t.Error("3,2,1,0,4 -> false")
	}
}
