package leetcode108

// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/

// TreeNode - Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Item to create TreeNode
type Item struct {
	Parent *TreeNode
	Start  int
	End    int
	IsLeft bool
}

func sortedArrayToBSTRecursive(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := start + (end-start)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBSTRecursive(nums, start, mid-1)
	root.Right = sortedArrayToBSTRecursive(nums, mid+1, end)
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return sortedArrayToBSTRecursive(nums, 0, len(nums)-1)
}

func sortedArrayToBSTIterative(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}

	dummy := &TreeNode{}
	item := Item{Parent: dummy, Start: 0, End: len(nums) - 1, IsLeft: true}
	stack := []Item{item}

	for len(stack) != 0 {
		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if it.Start <= it.End {
			mid := it.Start + (it.End-it.Start)/2
			ch := &TreeNode{Val: nums[mid]}
			if it.IsLeft {
				it.Parent.Left = ch
			} else {
				it.Parent.Right = ch
			}

			stack = append(stack, Item{Parent: ch, Start: it.Start, End: mid - 1, IsLeft: true})
			stack = append(stack, Item{Parent: ch, Start: mid + 1, End: it.End, IsLeft: false})
		}
	}

	return dummy.Left
}
