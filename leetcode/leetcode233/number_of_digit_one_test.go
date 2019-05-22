package leetcode233

import "testing"

func TestCountDigitOne(t *testing.T) {
	if countDigitOne(12) != 5 {
		t.Error("1 : 12 -> 5")
	}
	if countDigitOneMath(12) != 5 {
		t.Error("2 : 12 -> 5")
	}
}
