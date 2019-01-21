package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func removeLinkedListElements(head *ds.ListNode, val int) *ds.List {
	if head == nil {
		return head
	}

	p := head
	for p != nil {
		if p.Value == val {
			if p == head {
				p = p.Next
				head = p
			} else {
				if p.Next
			}
		}
	}
}