package sort

func radixSort(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	// find max
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	// temp slice
	temp := make([]int, 10)

	n := 1
	loop := max
	for loop > 0 {
		for _, v := range nums {
			for i := 0; i < 10; i++ {
				buckect := (v / n) % 10
				if buckect == i {
					temp[i]++
				}
			}
		}

		n *= 10
		loop /= 10
	}

	return
}
