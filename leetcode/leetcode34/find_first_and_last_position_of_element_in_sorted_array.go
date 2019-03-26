package leetcode34

// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	n := len(nums)
	start, end := -1, -1

	for i := 0; i < n; i++ {
		if nums[i] == target {
			start = i
			break
		}
	}

	for j := n - 1; j >= 0; j-- {
		if nums[j] == target {
			end = j
			break
		}
	}

	if start <= end {
		return []int{start, end}
	}
	return []int{-1, -1}
}

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

func searchRangeBinarySearch(nums []int, target int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{-1, -1}
	}

	start := binarySearchLastOne(nums, target-1) + 1
	end := binarySearchLastOne(nums, target)

	if end < len(nums) && start <= end {
		return []int{start, end}
	}

	return []int{-1, -1}
}
