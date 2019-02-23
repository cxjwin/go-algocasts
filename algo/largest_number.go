package algo

import "sort"
import "strconv"
import "strings"

// https://leetcode.com/problems/largest-number/

func compareNumber(a int, b int) bool {
	// number to string
	str1 := strconv.Itoa(a)
	str2 := strconv.Itoa(b)

	c1 := str1[0]
	c2 := str2[0]

	if c1 == c2 {
		s1 := str1 + str2
		s2 := str2 + str1
		return strings.Compare(s1, s2) > 0
	}

	if c1 > c2 {
		return true
	}
	return false
}

func largestNumber(nums []int) string {
	if nums == nil || len(nums) == 0 {
		return ""
	}

	sort.Slice(nums, func(i, j int) bool {
		return compareNumber(nums[i], nums[j])
	})

	if nums[0] == 0 {
		return "0"
	}

	str := ""
	for _, v := range nums {
		str = str + strconv.Itoa(v)
	}
	return str
}
