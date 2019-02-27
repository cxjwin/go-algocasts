package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

// https://leetcode.com/problems/partition-list/

func partitionList(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s := &ListNode{}
	g := &ListNode{}

	p := head
	ps := s
	pg := g

	for p != nil {
		if p.Val < x {
			ps.Next = p
			ps = ps.Next
		} else {
			pg.Next = p
			pg = pg.Next
		}
		p = p.Next
	}
	ps.Next = g.Next
	pg.Next = nil

	return s.Next
}
