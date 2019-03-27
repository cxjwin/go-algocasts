package leetcode55

// https://leetcode.com/problems/jump-game/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}

	n := len(nums)
	maxJ := 0

	for i := 0; i < n; i++ {
		if maxJ >= n-1 {
			return true
		}
		if maxJ < i {
			return false
		}
		maxJ = max(maxJ, i+nums[i])
	}

	return false
}
