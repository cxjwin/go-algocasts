package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func partitionList(head *ds.ListNode, x int) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s := &ds.ListNode{}
	g := &ds.ListNode{}

	p := head
	ps := s
	pg := g

	for p != nil {
		if p.Value.(int) < x {
			ps.Next = p
			ps = ps.Next
		} else {
			pg.Next = p
			pg = pg.Next
		}
		p = p.Next
	}
	ps.Next = g.Next

	return s.Next
}
