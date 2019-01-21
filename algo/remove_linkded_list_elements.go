package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func removeLinkedListElements(head *ds.ListNode, val int) *ds.ListNode {
	dummy := &ds.ListNode{Value: 0, Next: head}
	notEqual := dummy

	for notEqual.Next != nil {
		if notEqual.Next.Value == val {
			notEqual.Next = notEqual.Next.Next
		} else {
			notEqual = notEqual.Next
		}
	}

	return dummy.Next
}
