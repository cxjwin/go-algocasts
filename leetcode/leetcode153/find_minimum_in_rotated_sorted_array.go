package leetcode153

func findMin(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1

	for low < high {
		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}

func findMinEarlyReturn(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	low, high := 0, len(nums)-1

	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := low + (high-low)/2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return nums[low]
}
