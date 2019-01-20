package sort

// https://www.geeksforgeeks.org/radix-sort/

func radixSort(nums []int) {
	if nums == nil || len(nums) <= 1 {
		return
	}

	length := len(nums)

	// find max
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	countArr := make([]int, 10)      // count array
	outputArr := make([]int, length) // output array

	n := 1
	for max/n > 0 {

		// store count of occurrences in count array
		for _, v := range nums {
			buckect := (v / n) % 10
			countArr[buckect]++
		}

		// Change countArr[i] so that countArr[i] now contains actual
		// position of this digit in countArr[]
		for i := 1; i < 10; i++ {
			countArr[i] += countArr[i-1]
		}

		// build output
		for i := length - 1; i >= 0; i-- {
			buckect := (nums[i] / n) % 10
			outputArr[countArr[buckect]-1] = nums[i]
			countArr[buckect]--
		}

		for i, v := range outputArr {
			nums[i] = v
		}

		n *= 10
	}
}
