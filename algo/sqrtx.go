package algo

// https://leetcode.com/problems/sqrtx/

func sqrtBinarySearch(n int) int {
	low, high := int64(0), int64(n)

	for low <= high {
		mid := low + (high-low)/2
		mid2 := mid * mid

		if mid2 > int64(n) {
			high = mid - 1
		} else if mid < int64(n) {
			low = mid + 1
		} else {
			return int(mid)
		}
	}

	return int(high)
}

func sqrtNewton(n int) int {
	x := int64(n)
	for x*x > int64(n) {
		x = (x + int64(n)/x) / 2
	}
	return int(x)
}
