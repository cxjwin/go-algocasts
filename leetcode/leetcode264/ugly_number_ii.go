package leetcode264

// https://leetcode.com/problems/ugly-number-ii/

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a int, b int, c int) int {
	return min(min(a, b), c)
}

func nthUglyNumber(n int) int {
	if n <= 1 {
		return -1
	}

	p2, p3, p5 := 0, 0, 0
	uglyNums := make([]int, n)
	uglyNums[0] = 1

	for i := 1; i < n; i++ {
		min := min3(uglyNums[p2]*2, uglyNums[p3]*3, uglyNums[p5]*5)
		uglyNums[i] = min
		if min == uglyNums[p2]*2 {
			p2++
		}
		if min == uglyNums[p3]*3 {
			p3++
		}
		if min == uglyNums[p5]*5 {
			p5++
		}
	}

	return uglyNums[n-1]
}
