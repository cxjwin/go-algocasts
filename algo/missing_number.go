package algo

import "sort"

// https://leetcode.com/problems/missing-number/

func missingNumberUsingSum(nums []int) int {
	sum := 0

	for i, v := range nums {
		sum += i
		sum -= v
	}

	// add n, because total n + 1 numbers (0...n)
	sum += len(nums)

	if sum == len(nums) {
		return -1
	}

	return sum
}

func missingNumberUsingXOR(nums []int) int {
	res := 0

	// just like 'single number'
	for i, v := range nums {
		res ^= i
		res ^= v
	}

	// xor n, because total n + 1 numbers (0...n)
	res ^= len(nums)

	if res == len(nums) {
		return -1
	}

	return res
}

func missingNumberUsingSort(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}

	return -1
}
