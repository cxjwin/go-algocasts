package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func removeDuplicatesFromSortedList(head *ds.ListNode) *ds.ListNode {
	if head == nil {
		return head
	}

	cur := head
	next := cur.Next
	for next != nil {
		if cur.Value == next.Value {
			cur.Next = next.Next
		} else {
			cur = cur.Next
		}
		next = next.Next
	}

	return head
}
