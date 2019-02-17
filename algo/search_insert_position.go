package algo

func searchInsertPosition(nums []int, target int) int {
	if nums == nil {
		return -1
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + (high-low)/2
		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
