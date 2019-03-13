package leetcode264

import "testing"

func TestUglyNumber(t *testing.T) {
	res := nthUglyNumber(10)
	if res != 12 {
		t.Error("10th ugly number is 10")
	}
}
