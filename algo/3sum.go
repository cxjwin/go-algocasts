package algo

import "sort"

// https://leetcode.com/problems/3sum/

func threeSumToZeroOn2(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0)

	n := len(nums)

	for k := n - 1; k >= 2; k-- {
		if nums[k] < 0 {
			break
		}
		target := -nums[k]
		i, j := 0, k-1
		for i < j {
			if nums[i]+nums[j] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for i < j && nums[i+1] == nums[i] {
					i++
				}
				for i < j && nums[j-1] == nums[j] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else {
				j--
			}

		}
	}

	return res
}
