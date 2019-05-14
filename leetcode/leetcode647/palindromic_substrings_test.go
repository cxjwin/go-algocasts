package leetcode647

import "testing"

func TestCountSubstrings(t *testing.T) {
	s := "abc"
	res := countSubstrings(s)
	if res != 3 {
		t.Error("Three palindromic strings: a, b, c.")
	}

	s = "aaa"
	res = countSubstrings(s)
	if res != 6 {
		t.Error("Three palindromic strings: a, a, a, aa, aa, aaa.")
	}
}
