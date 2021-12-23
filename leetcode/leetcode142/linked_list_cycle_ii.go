package leetcode142

import (
	. "github.com/cxjwin/go-algocasts/datastructure"
)

// https://leetcode.com/problems/linked-list-cycle-ii/

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			for p := head; p != slow; p = p.Next {
				slow = slow.Next
			}
			return slow
		}
	}

	return nil
}
