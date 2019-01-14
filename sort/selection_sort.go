package sort

func selectionSort(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	l := len(nums)
	for i := 0; i < l-1; i++ {
		minIdx := i
		for j := i + 1; j < l; j++ {
			if nums[j] < nums[minIdx] {
				temp := nums[j]
				nums[j] = nums[minIdx]
				nums[minIdx] = temp
			}
		}
	}
}
