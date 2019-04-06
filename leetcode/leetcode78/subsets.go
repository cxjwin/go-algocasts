package leetcode78

import "math"

// https://leetcode.com/problems/subsets/

func _subsets(nums []int, start int, elem *[]int, res *[][]int) {
	// copy
	temp := make([]int, len(*elem))
	copy(temp, *elem)
	*res = append(*res, temp)

	for i := start; i < len(nums); i++ {
		*elem = append(*elem, nums[i])
		_subsets(nums, i+1, elem, res)
		*elem = (*elem)[:len(*elem)-1]
	}
}

func subsets(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}
	elem := &[]int{}
	res := &[][]int{}
	_subsets(nums, 0, elem, res)
	return *res
}

func subsetsBits(nums []int) [][]int {
	if nums == nil || len(nums) == 0 {
		return [][]int{}
	}

	n := len(nums)
	N := int(math.Pow(2, float64(n)))

	res := [][]int{}
	for i := 0; i < N; i++ {
		elem := []int{}
		for j := 0; j < n; j++ {
			if i>>uint(j)&1 == 1 {
				elem = append(elem, nums[j])
			}
		}
		res = append(res, elem)
	}

	return res
}
