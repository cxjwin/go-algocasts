package leetcode44

import "testing"

func TestIsMatch(t *testing.T) {
	if isMatch("aa", "a") {
		t.Error("aa is't match a")
	}
	if !isMatch("aa", "*") {
		t.Error("aa is't match *")
	}
	if isMatch("cb", "?a") {
		t.Error("cb is't match ?a")
	}
	if !isMatch("adceb", "*a*b") {
		t.Error("adceb is match *a*b")
	}
	if isMatch("acdcb", "a*c?b") {
		t.Error("acdcb is't match a*c?b")
	}
}
