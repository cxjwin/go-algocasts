package leetcode239

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	res := maxSlidingWindow(nums, 3)
	fmt.Println(res)
}
