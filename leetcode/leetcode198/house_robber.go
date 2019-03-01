package leetcode198

// https://leetcode.com/problems/house-robber/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	d := make([]int, len(nums))
	d[0] = nums[0]
	d[1] = max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		d[i] = max(d[i-1], d[i-2]+nums[i])
	}

	return d[len(nums)-1]
}

func robO1(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	prev1, prev2 := 0, 0

	for _, v := range nums {
		cur := max(prev1, prev2+v)
		prev2 = prev1
		prev1 = cur
	}

	return prev2
}
