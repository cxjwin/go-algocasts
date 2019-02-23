package algo

import "github.com/cxjwin/go-algocasts/utils"

// leetcode : https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}

	i := 0
	j := len(height) - 1
	max := 0
	for i < j {
		s := utils.IntMin(height[i], height[j]) * (j - i)
		max = utils.IntMax(s, max)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}
