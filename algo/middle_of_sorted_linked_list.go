package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func middleNodeTwoPass(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	count := 0
	p := head
	for p != nil {
		p = p.Next
		count++
	}

	mid := count / 2
	p = head
	for i := 0; i < mid; i++ {
		p = p.Next
	}

	return p
}

func middleNodeOnePass(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s, f := head, head

	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	return s
}
