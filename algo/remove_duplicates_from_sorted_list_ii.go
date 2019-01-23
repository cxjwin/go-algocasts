package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func removeDuplicatesFromSortedListII(head *ds.ListNode) *ds.ListNode {
	dummy := &ds.ListNode{Value: 0, Next: head}
	prev, cur := dummy, dummy.Next

	for cur != nil {
		for cur.Next != nil && cur.Value == cur.Next.Value {
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
