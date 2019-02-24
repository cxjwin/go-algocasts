package leetcode215

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	nums := []int{5, 3, 1, 8, 3, 2, 0}
	i := partition(nums, 0, len(nums)-1)
	fmt.Println(nums, i)
}

func TestKthLargestQuickSelect(t *testing.T) {
	nums := []int{4, 2, 8, 1, 8}
	res := findKthLargestQuickSelect(nums, 2)
	if res != 8 {
		t.Error("2th largest number is 8")
	}
}
