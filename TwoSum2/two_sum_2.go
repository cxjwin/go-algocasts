package main

import "fmt"

// Pair - two value struct
type Pair struct {
	firest, second interface{}
}

func getTwoNumSumToGivenValue(nums []int, target int) Pair {
	if len(nums) == 0 {
		return Pair{-1, -1}
	}

	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return Pair{i + 1, j + 1}
		}
	}

	return Pair{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	res := getTwoNumSumToGivenValue(nums, 9)
	fmt.Printf("find indexes : %v\n", res)
}
