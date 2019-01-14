package sort

import "github.com/cxjwin/go-algocasts/utils"

func partion1(nums []int, left int, right int) int {
	pivot := nums[left]
	i := left
	j := right

	for i != j {
		if nums[j] >= pivot && i < j {
			j--
		}

		if nums[i] <= pivot && i < j {
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

	p := partion1(nums, left, right)
	quickSortBody1(nums, left, p-1)
	quickSortBody1(nums, p+1, right)
}

func quickSort1(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	quickSortBody1(nums, 0, len(nums)-1)
}
