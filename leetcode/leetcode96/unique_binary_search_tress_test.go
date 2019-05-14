package leetcode96

import "testing"

func TestNumTrees(t *testing.T) {
	n := 3
	res := numTrees(n)
	if res != 5 {
		t.Error("Output: 5")
	}
}
