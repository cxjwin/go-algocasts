package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func mergeTwoSortedList(l *ds.ListNode, r *ds.ListNode) *ds.ListNode {
	dummy := &ds.ListNode{}
	p := dummy

	for l != nil && r != nil {
		if l.Value.(int) < r.Value.(int) {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}

	if l != nil {
		p.Next = l
	}

	if r != nil {
		p.Next = r
	}

	return dummy.Next
}
