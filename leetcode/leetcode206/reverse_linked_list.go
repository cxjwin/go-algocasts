package leetcode206

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/reverse-linked-list/

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	n := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return n
}

func reverseListIterative(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	cur := head
	var pre *ListNode

	for cur != nil {
		// next := cur.Next
		// cur.Next = pre
		// pre = cur
		// cur = next
		cur.Next, cur, pre = pre, cur.Next, cur
	}

	return pre
}
