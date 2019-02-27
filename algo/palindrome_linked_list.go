package algo

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
	"github.com/golang-collections/collections/stack"
)

// https://leetcode.com/problems/palindrome-linked-list/

func isPalindromeLinkedListUsingStack(list *List) bool {
	if list == nil || list.Len == 0 {
		return false
	}

	mid := list.Len / 2

	s := new(stack.Stack)

	node := list.Head
	idx := 0
	// iterate first half
	for idx < mid {
		s.Push(node.Val)
		node = node.Next
		idx++
	}

	// if odd number, skip mid number
	if list.Len%2 == 1 {
		node = node.Next
	}

	// iterate second half
	for node != nil {
		num := s.Pop()
		if num != node.Val {
			return false
		}
		node = node.Next
	}

	return true
}

func isPalindromeReverseLinkedList(list *List) bool {
	if list == nil || list.Len == 0 {
		return false
	}

	mid := list.Len / 2

	var pre *ListNode
	cur := list.Head
	idx := 0
	// iterate first half
	for idx < mid {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		idx++
	}

	// if odd number, skip mid number
	if list.Len%2 == 1 {
		cur = cur.Next
	}

	// iterate second half
	for pre != nil && cur != nil {
		if pre.Val != cur.Val {
			return false
		}
		pre = pre.Next
		cur = cur.Next
	}

	return true
}
