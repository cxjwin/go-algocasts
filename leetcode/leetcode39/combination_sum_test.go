package leetcode39

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	nums := []int{4, 2, 8}
	res := combinationSum(nums, 6)
	fmt.Println(res)
}
