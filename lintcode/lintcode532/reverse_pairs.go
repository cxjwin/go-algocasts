package lintcode532

func reversePairs(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	count := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] > nums[j] {
				count++
			}
		}
	}

	return count
}
