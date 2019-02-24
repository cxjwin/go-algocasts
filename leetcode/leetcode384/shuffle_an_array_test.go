package leetcode384

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	nums := []int{1, 2, 3}
	s := Constructor(nums)
	res := s.Shuffle()
	fmt.Println(res)

	res = s.Shuffle()
	fmt.Println(res)

	res = s.Shuffle()
	fmt.Println(res)
}
