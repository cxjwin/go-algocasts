package leetcode152

// https://leetcode.com/problems/maximum-product-subarray/

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

func max3(a int, b int, c int) int {
	return max(max(a, b), c)
}

func min3(a int, b int, c int) int {
	return min(min(a, b), c)
}

func maxProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	curMax, curMin, globalMax := nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		t1, t2 := nums[i]*curMax, nums[i]*curMin
		curMax = max3(nums[i], t1, t2)
		curMin = min3(nums[i], t1, t2)
		globalMax = max(curMax, globalMax)
	}

	return globalMax
}
