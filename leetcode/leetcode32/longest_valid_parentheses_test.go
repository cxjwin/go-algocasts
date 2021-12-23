package leetcode32

import "testing"

func TestLongestValidParentheses(t *testing.T) {
	type testFunc func(string) int

	testBody := func(f testFunc, t *testing.T) {
		s := ")()())"
		res := f(s)
		if res != 4 {
			t.Error("')()())' longest valid parentheses is 4.")
		}

		s = "(()"
		res = f(s)
		if res != 2 {
			t.Error("'(()' longest valid parentheses is 2.")
		}
	}

	testBody(longestValidParentheses, t)
	testBody(longestValidParenthesesDP, t)
}
