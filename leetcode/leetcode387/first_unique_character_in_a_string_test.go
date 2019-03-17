package leetcode387

import "testing"

func TestFirstUniqChar(t *testing.T) {
	s := "leetcode"
	res := firstUniqChar(s)
	if res != 0 {
		t.Error("leetcode -> 0")
	}
	s = "loveleetcode"
	res = firstUniqChar(s)
	if res != 2 {
		t.Error("loveleetcode -> 2")
	}
}

func TestFirstUniqCharOnePass(t *testing.T) {
	s := "leetcode"
	res := firstUniqCharOnePass(s)
	if res != 0 {
		t.Error("leetcode -> 0")
	}
	s = "loveleetcode"
	res = firstUniqCharOnePass(s)
	if res != 2 {
		t.Error("loveleetcode -> 2")
	}
}
