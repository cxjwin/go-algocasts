package leetcode771

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	J := "aA"
	S := "aAAbbbb"
	res := numJewelsInStones(J, S)
	if res != 3 {
		t.Error("aA <-> aAAbbbb is 3")
	}
}
