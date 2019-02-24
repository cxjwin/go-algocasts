package leetcode215

// https://leetcode.com/problems/kth-largest-element-in-an-array/

func swap(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func partition(nums []int, low int, high int) int {
	pivot := nums[low]
	i, j := low, high
	for i != j {
		for i < j && nums[j] <= pivot {
			j--
		}
		for i < j && nums[i] >= pivot {
			i++
		}
		if i < j {
			swap(nums, i, j)
		}
	}
	nums[low] = nums[i]
	nums[i] = pivot

	return i
}

func findKthLargestQuickSelect(nums []int, k int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		p := partition(nums, low, high)
		if p == k-1 {
			return nums[p]
		}
		if p > k-1 {
			high = p - 1
		} else {
			low = p + 1
		}
	}

	return -1
}
