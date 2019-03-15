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

func longestConsecutiveWithMap(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	m := make(map[int]int)

	maxL := 0

	for _, v := range nums {
		m[v] = v
	}

	n := len(nums)

	for i := 0; i < n && len(m) != 0; i++ {
		delete(m, nums[i])
		low, high := nums[i], nums[i]

		for true {
			low--
			_, ok := m[low]
			if ok {
				delete(m, low)
			} else {
				break
			}
		}

		for true {
			high++
			_, ok := m[high]
			if ok {
				delete(m, high)
			} else {
				break
			}
		}

		maxL = max(maxL, high-low-1)
	}

	return maxL
}
