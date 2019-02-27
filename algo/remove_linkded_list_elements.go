package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

func removeLinkedListElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	notEqual := dummy

	for notEqual.Next != nil {
		if notEqual.Next.Val == val {
			notEqual.Next = notEqual.Next.Next
		} else {
			notEqual = notEqual.Next
		}
	}

	return dummy.Next
}
