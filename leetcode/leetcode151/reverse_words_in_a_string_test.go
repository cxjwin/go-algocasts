package leetcode151

import "testing"

func TestReverseWords(t *testing.T) {
	case1 := "the sky is blue"
	res := reverseWords(case1)
	if res != "blue is sky the" {
		t.Error("'the sky is blue' -> 'blue is sky the'")
	}

	case2 := "  hello world!  "
	res = reverseWords(case2)
	if res != "world! hello" {
		t.Error("'  hello world!  ' -> 'world! hello'")
	}

	case3 := "a good   example"
	res = reverseWords(case3)
	if res != "example good a" {
		t.Error("'a good   example' -> 'example good a'")
	}

	case4 := " "
	res = reverseWords(case4)
	if res != "" {
		t.Error("' ' -> ''")
	}
}
