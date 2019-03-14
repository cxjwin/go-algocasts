package leetcode239

// https://leetcode.com/problems/sliding-window-maximum/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil || len(nums) == 0 {
		return nums
	}

	n := len(nums)
	res := make([]int, n-k+1)

	maxFromLeft := make([]int, n)
	maxFromRight := make([]int, n)
	maxFromLeft[0] = nums[0]
	maxFromRight[n-1] = nums[n-1]

	for i, j := 1, n-2; i < n; i, j = i+1, j-1 {
		if i%k == 0 {
			maxFromLeft[i] = nums[i]
		} else {
			maxFromLeft[i] = max(maxFromLeft[i-1], nums[i])
		}

		if j%k == k-1 {
			maxFromRight[j] = nums[j]
		} else {
			maxFromRight[j] = max(maxFromRight[j+1], nums[j])
		}
	}

	for i := 0; i <= n-k; i++ {
		res[i] = max(maxFromRight[i], maxFromLeft[i+k-1])
	}

	return res
}
