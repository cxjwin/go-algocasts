package leetcode338

// https://leetcode.com/problems/counting-bits/

func countBitsOfNumber(num int) int {
	count := 0
	for num != 0 {
		count++
		num &= (num - 1)
	}
	return count
}

func countBits(num int) []int {
	d := make([]int, num+1)
	for i := 0; i <= num; i++ {
		d[i] = countBitsOfNumber(i)
	}
	return d
}

func countBitsDP(num int) []int {
	d := make([]int, num+1)

	for i := 1; i <= num; i++ {
		d[i] = d[i&(i-1)] + 1
	}

	return d
}
