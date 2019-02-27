package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

func removeNthNodeFromEndOfList(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
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
