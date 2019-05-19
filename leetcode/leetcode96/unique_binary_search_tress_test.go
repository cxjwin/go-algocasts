package leetcode96

import "testing"

func TestNumTrees(t *testing.T) {
	if numTrees(3) != 5 {
		t.Error("1 : Output: 5")
	}
	if numTreesOn(3) != 5 {
		t.Error("2 : Output: 5")
	}
}
