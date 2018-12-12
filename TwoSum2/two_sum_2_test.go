package algo

import (
	"fmt"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := getTwoNumSumToGivenValue(nums, 9)
	fmt.Printf("find indexes : %v\n", res)
	if res.first != 1 || res.second != 2 {
		t.Error("res is {1, 2}")
	}
}
