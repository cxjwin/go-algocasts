package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func rotateListRight(head *ds.ListNode, k int) *ds.ListNode {
	if head == nil || head.Next == nil || k <= 0 {
		return head
	}

	p := head

	count := 0
	last := p
	for p != nil {
		count++
		if p.Next == nil {
			last = p
		}
		p = p.Next
	}
	last.Next = head

	m := k % count
	n := count - m

	s := head
	for i := 0; i < n-1; i++ {
		s = s.Next
	}
	newHead := s.Next
	s.Next = nil

	return newHead
}
