package leetcode581

import "sort"

// https://leetcode.com/problems/shortest-unsorted-continuous-subarray/

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	copyNums := make([]int, n)
	copy(copyNums, nums)
	sort.Ints(copyNums)

	start, end := -1, -1

	for i := 0; i < n; i++ {
		if copyNums[i] != nums[i] {
			start = i
			break
		}
	}

	for j := n - 1; j >= 0; j-- {
		if copyNums[j] != nums[j] {
			end = j
			break
		}
	}

	if start != -1 && end != -1 {
		return end - start + 1
	}

	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func findUnsortedSubarrayTwoPass(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)

	i, j := 0, n-1
	for i < n-1 && nums[i] <= nums[i+1] {
		i++
	}
	for j > 0 && nums[j] >= nums[j-1] {
		j--
	}
	// min & max
	minN, maxN := 1<<31-1, -1<<31
	for k := i; k <= j; k++ {
		minN = min(minN, nums[k])
		maxN = max(maxN, nums[k])
	}

	// expand
	for i >= 0 && nums[i] > minN {
		i--
	}
	for j < n && nums[j] < maxN {
		j++
	}

	return max(0, j-i-1)
}

func findUnsortedSubarrayOnePass(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)

	i, j := 0, -1

	minN, maxN := 1<<31-1, -1<<31

	for k := 0; k < n; k++ {
		maxN = max(maxN, nums[k])
		if nums[k] != maxN {
			j = k
		}
		p := n - 1 - k
		minN = min(minN, nums[p])
		if nums[p] != minN {
			i = p
		}
	}

	return j - i + 1
}
