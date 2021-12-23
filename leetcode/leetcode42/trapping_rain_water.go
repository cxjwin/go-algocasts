package leetcode42

// https://leetcode.com/problems/trapping-rain-water/

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	leftMax, rightMax := -1, -1
	d := make([]int, len(height))

	for i := 0; i < len(height); i++ {
		d[i] = max(leftMax, height[i])
		leftMax = d[i]
	}

	sum := 0
	for j := len(height) - 1; j >= 0; j-- {
		rightMax = max(rightMax, height[j])
		d[j] = min(d[j], rightMax)
		sum += d[j] - height[j]
	}
	return sum
}

func trapO1(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	leftMax, rightMax, sum := -1, -1, 0

	i, j := 0, len(height)-1

	for i <= j {
		leftMax = max(leftMax, height[i])
		rightMax = max(rightMax, height[j])
		if leftMax < rightMax {
			sum += leftMax - height[i]
			i++
		} else {
			sum += rightMax - height[j]
			j--
		}
	}

	return sum
}
