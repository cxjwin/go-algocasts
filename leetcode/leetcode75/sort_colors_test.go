package leetcode75

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors2(nums)
	fmt.Println(nums)
}
