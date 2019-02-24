package algo

// https://leetcode.com/problems/hamming-distance/

func numberOfOne(n int) int {
	count := 0
	for n != 0 {
		count++
		n &= (n - 1)
	}
	return count
}

func hammingDistance(x int, y int) int {
	return numberOfOne(x ^ y)
}
