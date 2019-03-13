package leetcode128

import "sort"

// https://leetcode.com/problems/longest-consecutive-sequence/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// sort
	sort.Ints(nums)

	n := len(nums)
	m := 1

	for i := 0; i < n; i++ {
		len := 1
		for i < n-1 {
			if nums[i+1]-nums[i] > 1 {
				break
			}

			if nums[i+1]-nums[i] == 1 {
				len++
			}
			i++
		}
		m = max(m, len)
	}

	return m
}
