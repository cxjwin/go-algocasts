package leetcode518

import "testing"

func TestChange(t *testing.T) {
	coins := []int{1, 2, 5}
	res := change(5, coins)
	if res != 4 {
		t.Error("[1, 2, 5] <-> 5 => 5")
	}
}

func TestChangeDP(t *testing.T) {
	coins := []int{1, 2, 5}
	res := changeDP(5, coins)
	if res != 4 {
		t.Error("[1, 2, 5] <-> 5 => 5")
	}
}
