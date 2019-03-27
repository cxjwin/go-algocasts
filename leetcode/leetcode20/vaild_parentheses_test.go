package leetcode20

import "testing"

func TestIsValid(t *testing.T) {
	case1 := "()"
	if !isValid(case1) {
		t.Error("() is valid")
	}
	case2 := "()[]{}"
	if !isValid(case2) {
		t.Error("()[]{} is vaild")
	}
	case3 := "(]"
	if isValid(case3) {
		t.Error("(] is not vaild")
	}
	case4 := "([)]"
	if isValid(case4) {
		t.Error("([)] is not vaild")
	}
	case5 := "{[]}"
	if !isValid(case5) {
		t.Error("{[]} is vaild")
	}
	case6 := "]"
	if isValid(case6) {
		t.Error("] is not vaild")
	}
}
