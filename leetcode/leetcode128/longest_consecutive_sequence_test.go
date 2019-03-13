package leetcode128

import "testing"

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	n := longestConsecutive(nums)
	if n != 4 {
		t.Error("The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.")
	}

	nums = []int{1, 2, 0, 1}
	n = longestConsecutive(nums)
	if n != 3 {
		t.Error("The longest consecutive elements sequence is [1, 1, 2]. Therefore its length is 3.")
	}
}
