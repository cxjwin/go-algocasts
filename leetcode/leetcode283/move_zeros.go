package leetcode283

// https://leetcode.com/problems/move-zeroes/

func moveZeroes(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)

	i := 0
	end := n - 1
	for i <= end {
		if nums[i] == 0 {
			j := i + 1
			for j <= end {
				nums[j-1] = nums[j]
				j++
			}
			nums[end] = 0
			end--
		} else {
			i++
		}
	}
}

func moveZerosAssign(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)

	slow, fast := 0, 0

	for ; fast < n; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}

	for ; slow < n; slow++ {
		nums[slow] = 0
	}
}

func moveZerosSwap(nums []int) {
	if nums == nil || len(nums) == 0 {
		return
	}

	n := len(nums)
	for slow, fast := 0, 0; fast < n; fast++ {
		if nums[fast] != 0 {
			temp := nums[slow]
			nums[slow] = nums[fast]
			nums[fast] = temp
			slow++
		}
	}
}
