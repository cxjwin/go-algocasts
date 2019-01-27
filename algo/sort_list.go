package algo

import "github.com/cxjwin/go-algocasts/datastructure"

func swapListNode(l *ds.ListNode, r *ds.ListNode) {
	temp := l.Value
	l.Value = r.Value
	r.Value = temp
}

func quickSort(head *ds.ListNode, end *ds.ListNode) {
	if head == end || head.Next == end {
		return
	}

	pivot := head.Value
	s := head
	f := s.Next

	for f != end {
		if f.Value.(int) <= pivot.(int) {
			s = s.Next
			swapListNode(s, f)
		}
		f = f.Next
	}
	swapListNode(head, s)

	quickSort(head, s)
	quickSort(s.Next, end)
}

func quickSortList(head *ds.ListNode) *ds.ListNode {
	quickSort(head, nil)
	return head
}

func mergeSortList(head *ds.ListNode) *ds.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	s, f := head, head

	for f.Next != nil && f.Next.Next != nil {
		s = s.Next
		f = f.Next.Next
	}

	r := mergeSortList(s.Next)
	s.Next = nil
	l := mergeSortList(head)
	return mergeTwoSortedList(l, r)
}
