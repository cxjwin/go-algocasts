package leetcode322

import "testing"

func TestCoinChangeDP(t *testing.T) {
	coins := []int{1, 2, 5}
	amount := 11
	res := coinChangeDP(coins, amount)
	if res != 3 {
		t.Error("[1, 2, 5] <-> 11 => 3")
	}
}
