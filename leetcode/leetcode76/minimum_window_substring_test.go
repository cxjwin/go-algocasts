package leetcode76

import "testing"

func TestMinWindow(ts *testing.T) {
	s, t := "ADOBECODEBANC", "ABC"

	res := minWindow(s, t)

	if res != "BANC" {
		ts.Error("BANC")
	}
}
