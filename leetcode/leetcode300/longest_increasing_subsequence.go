package leetcode300

// https://leetcode.com/problems/longest-increasing-subsequence/

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	n := len(nums)
	d := make([]int, n)
	d[0] = 1
	m := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			var cur int
			if nums[j] < nums[i] {
				cur = d[j] + 1
			} else {
				cur = 1
			}
			d[i] = max(cur, d[i])
		}
		m = max(m, d[i])
	}

	return d[n-1]
}

func binarySearch(nums []int, l int, x int) int {
	low, high := 0, l-1

	for low <= high {
		mid := low + (high-low)/2
		if x > nums[mid] {
			low = mid + 1
		} else if x < nums[mid] {
			high = mid - 1
		} else {
			return mid
		}
	}

	return low
}

func lengthOfLISBinarySearch(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	d := make([]int, len(nums))
	l := 0

	for _, v := range nums {
		i := binarySearch(d, l, v)
		d[i] = v
		if i == l {
			l++
		}
	}

	return l
}
