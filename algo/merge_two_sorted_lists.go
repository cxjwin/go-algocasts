package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/merge-two-sorted-lists/

func mergeTwoSortedList(l *ListNode, r *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: nil}
	p := dummy

	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}

	if l != nil {
		p.Next = l
	}

	if r != nil {
		p.Next = r
	}

	return dummy.Next
}
