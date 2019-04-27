package leetcode45

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func jump(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	n := len(nums)
	maxLen := 0
	d := make([]int, n)

	for i := 0; i < n; i++ {
		if maxLen >= n-1 {
			return d[n-1]
		}
		if i > maxLen {
			return -1
		}

		maxLen = max(i+nums[i], maxLen)
		last := min(maxLen, n-1)
		for j := last; j > i && d[j] == 0; j-- {
			d[j] = d[i] + 1
		}
	}

	return -1
}

func jumpO1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	n, maxLen, jumps, curEnd := len(nums), 0, 0, 0

	for i := 0; i < n; i++ {
		if maxLen >= n-1 {
			return jumps + 1
		}
		if i > maxLen {
			return -1
		}

		if i > curEnd {
			jumps++
			curEnd = maxLen
		}

		maxLen = max(maxLen, i+nums[i])
	}

	return -1
}
