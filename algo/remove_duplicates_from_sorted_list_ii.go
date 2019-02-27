package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

func removeDuplicatesFromSortedListII(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	prev, cur := dummy, dummy.Next

	for cur != nil {
		for cur.Next != nil && cur.Val == cur.Next.Val {
			cur = cur.Next
		}
		if prev.Next != cur {
			prev.Next = cur.Next
		} else {
			prev = prev.Next
		}
		cur = prev.Next
	}

	return dummy.Next
}
