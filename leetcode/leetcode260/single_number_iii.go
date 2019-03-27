package leetcode260

// https://leetcode.com/problems/single-number-iii/

func singleNumber(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	mask := 1
	for (xor & mask) == 0 {
		mask <<= 1
	}
	res := []int{0, 0}
	for _, v := range nums {
		if mask&v == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}

func singleNumber2(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}
	mask := xor & (-xor)

	res := []int{0, 0}
	for _, v := range nums {
		if mask&v == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}
