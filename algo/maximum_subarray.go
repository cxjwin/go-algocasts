package algo

import "math"
import "github.com/cxjwin/go-algocasts/utils"

func maxSubArray(nums []int) int {
	max := int(math.MinInt32)
	cur := 0
	for i := 0; i < len(nums); i++ {
		if cur <= 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		max = utils.IntMax(max, cur)
	}
	return max
}
