package leetcode78

import (
	"fmt"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsets(nums)
	fmt.Println(res)
	if len(res) != 8 {
		t.Error("output len is 8")
	}
}

func TestSubsetsBits(t *testing.T) {
	nums := []int{1, 2, 3}
	res := subsetsBits(nums)
	fmt.Println(res)
	if len(res) != 8 {
		t.Error("output len is 8")
	}
}
