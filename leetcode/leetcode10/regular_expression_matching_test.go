package leetcode10

import "testing"

func TestIsMatch(t *testing.T) {
	if isMatch("aa", "a") {
		t.Error("aa is't match a")
	}
	if !isMatch("aa", "a*") {
		t.Error("aa is match a*")
	}
	if !isMatch("aa", ".*") {
		t.Error("aa is match .*")
	}
	if !isMatch("aab", "c*a*b*") {
		t.Error("aab is match c*a*b*")
	}
	if isMatch("mississippi", "mis*is*p*.") {
		t.Error("mississippi is't match mis*is*p*.")
	}
}
