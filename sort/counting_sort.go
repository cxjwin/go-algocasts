package sort

// https://www.geeksforgeeks.org/counting-sort/

func countingSort(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	min := nums[0]
	max := nums[0]

	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	k := max - min + 1
	temp := make([]int, k)

	for _, v := range nums {
		temp[v-min]++
	}

	idx := 0
	for i, v := range temp {
		count := v
		for count > 0 {
			nums[idx] = i + min
			idx++
			count--
		}
	}
}
