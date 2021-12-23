package leetcode24

import . "cxjwin.com/go-algocasts/datastructure"

// https://leetcode.com/problems/swap-nodes-in-pairs/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fir := head
	sec := fir.Next
	head = sec

	for fir != nil && sec != nil {
		fNext := fir.Next.Next
		var sNext *ListNode
		if fNext != nil {
			sNext = fNext.Next
		}

		sec.Next = fir
		if sNext != nil {
			fir.Next = sNext
		} else {
			fir.Next = fNext
		}

		fir = fNext
		sec = sNext
	}

	if fir != nil && sec == nil {
		fir.Next = nil
	}

	return head
}
