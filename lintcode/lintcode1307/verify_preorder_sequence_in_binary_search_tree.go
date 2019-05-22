package lintcode1307

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://www.lintcode.com/problem/verify-preorder-sequence-in-binary-search-tree/

func verify(nums []int, start int, end int) bool {
	if start == end || start+1 == end {
		return true
	}

	root := nums[start]
	i := start + 1

	for i < end && nums[i] < root {
		i++
	}
	mid := i
	for i < end && nums[i] > root {
		i++
	}
	if i == end {
		return verify(nums, start+1, mid) && verify(nums, mid, end)
	}

	return false
}

func verifyPreorderDivideConquer(nums []int) bool {
	if nums == nil {
		return false
	}
	return verify(nums, 0, len(nums))
}

func verifyPreorderStack(nums []int) bool {
	if nums == nil {
		return false
	}

	stack := make(Stack, 0)
	lowerBound := (-1 << 31)
	for _, v := range nums {
		if v < lowerBound {
			return false
		}

		for !stack.Empty() && v > stack.Top() {
			lowerBound = stack.Pop()
		}

		stack.Push(v)
	}

	return true
}
