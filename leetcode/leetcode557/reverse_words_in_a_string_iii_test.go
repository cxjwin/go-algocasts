package leetcode557

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	s := "Let's take LeetCode contest"
	s = reverseWords(s)
	fmt.Println(s)
}
