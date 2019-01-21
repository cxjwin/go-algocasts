package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func intersectionOfTwoLinkedLists(headA, headB *ds.ListNode) *ds.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	p, q := headA, headB

	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}

		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}

	return p
}

func intersectionOfTwoLinkedListsWithLength(headA, headB *ds.ListNode) *ds.ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lenA := 0
	p := headA
	for p != nil {
		lenA++
		p = p.Next
	}

	lenB := 0
	q := headB
	for q != nil {
		lenB++
		q = q.Next
	}

	p = headA
	q = headB

	delta := lenA - lenB
	len := 0
	if delta > 0 {
		for i := 0; i < delta; i++ {
			p = p.Next
		}
		len = lenB
	} else {
		delta = -delta
		for i := 0; i < delta; i++ {
			q = q.Next
		}
		len = lenA
	}

	for i := 0; i < len; i++ {
		if p != q {
			p = p.Next
			q = q.Next
		}
	}

	return p
}
