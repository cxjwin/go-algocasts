package leetcode34

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func binarySearchLastOne(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)/2
		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return high
}

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	end := binarySearchLastOne(nums, target)
	start := binarySearchLastOne(nums, target-1) + 1
	if start >= 0 && start <= end && end < len(nums) {
		return []int{start, end}
	}
	return []int{-1, -1}
}
