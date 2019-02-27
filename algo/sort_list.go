package algo

import . "github.com/cxjwin/go-algocasts/datastructure"

func swapListNode(l *ListNode, r *ListNode) {
	temp := l.Val
	l.Val = r.Val
	r.Val = temp
}

func quickSort(head *ListNode, end *ListNode) {
	if head == end || head.Next == end {
		return
	}

	pivot := head.Val
	s := head
	f := s.Next

	for f != end {
		if f.Val <= pivot {
			s = s.Next
			swapListNode(s, f)
		}
		f = f.Next
	}
	swapListNode(head, s)

	quickSort(head, s)
	quickSort(s.Next, end)
}

func quickSortList(head *ListNode) *ListNode {
	quickSort(head, nil)
	return head
}

func mergeSortList(head *ListNode) *ListNode {
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
