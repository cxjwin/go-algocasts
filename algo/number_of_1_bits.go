package algo

// https://leetcode.com/problems/number-of-1-bits/

func hammingWeightWithMask(num uint32) int {
	var mask uint32 = 1
	count := 0
	for mask != 0 {
		if num&mask != 0 {
			count++
		}
		mask <<= 1
	}
	return count
}

func hammingWeight(num uint32) int {
	count := 0

	for num != 0 {
		count++
		num &= (num - 1)
	}

	return count
}
