package algo

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	res := getTwoNumSumToGivenValueHashMap(nums, 9)
	fmt.Printf("find indexes : %v\n", res)

	res = getTwoNumSumToGivenValueHashMap(nums, 14)
	fmt.Printf("find indexes : %v\n", res)
}
