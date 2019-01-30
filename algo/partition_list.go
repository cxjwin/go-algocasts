package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func partitionList(head *ds.ListNode, x int) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ds.ListNode{}
	s := dummy
	g := &ds.ListNode{}
	p := head

	for p != nil {
		if p.Value.(int) < x {
			s.Next = p
			s = s.Next
		} else {
			g.Next = p
			g = g.Next
		}
		p = p.Next
	}

	return dummy.Next
}
