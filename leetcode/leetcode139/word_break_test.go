package leetcode139

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	m := map[string]bool{"1": true}
	num := 1
	res := false

	if _, ok := m["1"]; ok && num == 1 {

	}

	fmt.Println(res)
}

func TestSubstring(t *testing.T) {
	s := "leetcode"
	fmt.Println(s[0:1])
}

func TestWordBreak(t *testing.T) {
	s := "leetcode"
	wordDict := []string{"leet", "code"}
	res := wordBreak(s, wordDict)
	if !res {
		t.Error("Return true because 'leetcode' can be segmented as 'leet code'.")
	}
}
