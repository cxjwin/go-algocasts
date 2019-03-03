package leetcode213

// https://leetcode.com/problems/house-robber-ii/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func _rob(nums []int, start int, end int) int {
	prev2, prev1 := 0, 0
	for i := start; i <= end; i++ {
		cur := max(prev1, prev2+nums[i])
		prev2 = prev1
		prev1 = cur
	}
	return prev1
}

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(_rob(nums, 0, len(nums)-2), _rob(nums, 1, len(nums)-1))
}
