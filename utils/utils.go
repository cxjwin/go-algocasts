package utils

// IntMax max func with int type
func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// IntMin min func with int type
func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Swap in int slice
func Swap(nums []int, i int, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}
