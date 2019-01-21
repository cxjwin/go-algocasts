package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func removeNthNodeFromEndOfList(head *ds.ListNode, n int) *ds.ListNode {
	dummy := &ds.ListNode{Value: 0, Next: head}
	p, q := dummy, dummy

	i := n
	for ; i > 0 && p.Next != nil; i-- {
		p = p.Next
	}
	if i != 0 {
		return dummy.Next
	}

	for p.Next != nil {
		p = p.Next
		q = q.Next
	}

	q.Next = q.Next.Next

	return dummy.Next
}
