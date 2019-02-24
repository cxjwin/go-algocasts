package algo

// https://leetcode.com/problems/single-element-in-a-sorted-array/

func singleNonDuplicateWithXOR(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func singleNonDuplicateBinarySearch(nums []int) int {
	n := len(nums)
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		if mid-1 >= 0 && nums[mid-1] == nums[mid] {
			mid--
		} else if mid+1 < n && nums[mid+1] == nums[mid] {
			//
		} else {
			return nums[mid]
		}
		if (mid-low)&1 == 1 {
			high = mid - 1
		} else {
			low = mid + 2
		}
	}
	return -1
}
