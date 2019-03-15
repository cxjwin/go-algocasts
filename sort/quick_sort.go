package sort

// https://www.geeksforgeeks.org/quick-sort/

import "github.com/cxjwin/go-algocasts/utils"

func partion1(nums []int, left int, right int) int {
	pivot := nums[left]
	i := left
	j := right

	for i != j {
		for nums[j] >= pivot && i < j {
			j--
		}

		for nums[i] <= pivot && i < j {
			i++
		}

		if i < j {
			utils.Swap(nums, i, j)
		}
	}

	nums[left] = nums[i]
	nums[i] = pivot

	return i
}

func quickSortBody1(nums []int, left int, right int) {
	if left >= right {
		return
	}

	i := partion1(nums, left, right)
	quickSortBody1(nums, left, i-1)
	quickSortBody1(nums, i+1, right)
}

func quickSort1(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	quickSortBody1(nums, 0, len(nums)-1)
}

func partion2(nums []int, left int, right int) int {
	pivot := nums[right]
	i := left
	j := right
	for i != j {
		for nums[i] <= pivot && i < j {
			i++
		}

		for nums[j] >= pivot && i < j {
			j--
		}

		if i < j {
			utils.Swap(nums, i, j)
		}
	}

	nums[right] = nums[i]
	nums[i] = pivot

	return i
}

func quickSortBody2(nums []int, left int, right int) {
	if left >= right {
		return
	}

	i := partion2(nums, left, right)
	quickSortBody2(nums, left, i-1)
	quickSortBody2(nums, i+1, right)
}

func quickSort2(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	quickSortBody2(nums, 0, len(nums)-1)
}

func partion3(nums []int, left int, right int, pivotIndex int) int {
	if pivotIndex != right {
		utils.Swap(nums, pivotIndex, right)
	}

	pivot := nums[right]
	i := left

	// from left to right-1
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			utils.Swap(nums, i, j)
			i++
		}
	}
	utils.Swap(nums, i, right)
	return i
}

func quickSortBody3(nums []int, left int, right int) {
	if left >= right {
		return
	}

	i := partion3(nums, left, right, (left+right)/2)
	quickSortBody3(nums, left, i-1)
	quickSortBody3(nums, i+1, right)
}

func quickSort3(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	quickSortBody3(nums, 0, len(nums)-1)
}
