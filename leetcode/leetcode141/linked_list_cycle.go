package leetcode141

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/linked-list-cycle/

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}

	return false
}
